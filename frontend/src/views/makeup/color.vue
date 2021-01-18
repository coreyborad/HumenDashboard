<template>
  <div>
    <el-dialog width="80%" title="色號詳細資訊" v-if="visible" :visible="true" :close-on-click-modal="false" @close="close">
      <!-- <div class="tooltips">
      <el-button type="primary" size="small" @click="stockDialogVisible = true">新增持有股票</el-button>
    </div> -->
      <el-table
        v-loading="loading"
        :data="colorList"
        border
        fit
        highlight-current-row
      >
        <el-table-column align="center" label="品項">
          <template slot-scope="scope">
            {{ scope.row.name }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="色號">
          <template slot-scope="scope">
            <!-- <el-popover
              placement="right"
              width="400"
              trigger="click">
              test
              <span slot="reference"> {{ scope.row.color_name }}</span>
            </el-popover> -->
            {{ scope.row.color_name }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="總成本" width="100">
          <template slot-scope="scope">
            {{ $_toCurrency(scope.row.cost_total) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="總銷售" width="100">
          <template slot-scope="scope">
            {{ $_toCurrency(scope.row.sale_total) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="存貨狀況(剩餘/總共)" width="150">
          <template slot-scope="scope">
            {{ scope.row.inventory }} / {{ scope.row.total_cost }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作">
          <template slot-scope="scope">
            <el-button size="small" @click="showCost(scope.row)">成本列表</el-button>
            <el-button size="small" @click="showSale(scope.row)">銷售列表</el-button>
            <el-popconfirm
              title="確定刪除?"
              @onConfirm="deleteColor(scope.row)"
            >
              <el-button slot="reference" style="margin-left: 8px" type="danger" icon="el-icon-delete" size="small" />
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <!-- <div slot="footer" class="dialog-footer">
        <el-button @click="close">取消</el-button>
      </div> -->
    </el-dialog>
    <CostDialog ref="costDialog" :visible.sync="costDialogVisible" />
    <SaleDialog ref="saleDialog" :visible.sync="saleDialogVisible" />
  </div>
</template>

<script>

import CostDialog from './cost.vue'
import SaleDialog from './sale.vue'

export default {
  components: {
    CostDialog,
    SaleDialog
  },
  props: {
    visible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      loading: false,
      costDialogVisible: false,
      saleDialogVisible: false,
      itemDialogVisible: false
    }
  },
  computed: {
    colorList() {
      const list = this.$store.getters['makeup/colorList']
      return list.map(l => {
        // 成本狀況
        l.cost_total = 0
        if (l.costs.length > 0) {
          l.cost_total = l.costs
            .map(el => el.price * el.count)
            .reduce((a, b) => { return a + b })
        }
        // 銷售狀況
        l.sale_total = 0
        if (l.sales.length > 0) {
          l.sale_total = l.sales
            .map(el => el.price * el.count)
            .reduce((a, b) => { return a + b })
        }
        // 存貨狀況
        let cost_count = 0
        let sale_count = 0
        if (l.costs.length > 0) {
          cost_count = l.costs
            .map(el => el.count)
            .reduce((a, b) => { return a + b })
        }
        if (l.sales.length > 0) {
          sale_count = l.sales
            .map(el => el.count)
            .reduce((a, b) => { return a + b })
        }
        l.total_cost = cost_count
        l.inventory = cost_count - sale_count
        return l
      })
    }
  },
  async created() {
  },
  methods: {
    async deleteColor(color) {
      this.loading = true
      try {
        const colorList = await this.$store.dispatch('makeup/deleteMakeupInfo', color.id)
        this.$message.success('刪除成功')
        if (this.colorList.length <= 1) {
          this.$emit('update:visible', false)
        }
      } catch (error) {
        this.$message.error(error)
      }
      this.loading = false
    },
    async showCost(color) {
      await this.$store.dispatch('makeup/setCurrentValue', { target: 'color_name', value: color.color_name })
      await this.$store.dispatch('makeup/setCurrentValue', { target: 'id', value: color.id })
      this.costDialogVisible = true
    },
    async showSale(color) {
      await this.$store.dispatch('makeup/setCurrentValue', { target: 'color_name', value: color.color_name })
      await this.$store.dispatch('makeup/setCurrentValue', { target: 'id', value: color.id })
      this.saleDialogVisible = true
    },
    close() {
      this.$emit('update:visible', false)
    }
  }
}
</script>
