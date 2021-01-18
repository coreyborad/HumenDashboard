<template>
  <div>
    <el-dialog title="成本列表" v-if="visible" :visible="true" :close-on-click-modal="false" @close="close">
      <div class="tooltips">
        <el-button type="primary" size="small" @click="createCost">新增進貨</el-button>
      </div>
      <el-table
        v-loading="loading"
        :data="costList"
        border
        fit
        highlight-current-row
      >
        <el-table-column align="center" label="成本" width="160">
          <template slot-scope="scope">
            {{ $_toCurrency(scope.row.price) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="數量" width="80">
          <template slot-scope="scope">
            {{ scope.row.count }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="進貨日期">
          <template slot-scope="scope">
            {{ scope.row.order_date }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作">
          <template slot-scope="scope">
            <el-popover
              placement="right"
              width="400"
              trigger="click"
              @show="setUpdateData(scope.row)"
              >
              <CostForm ref="form" :form.sync="form" />
              <el-button type="primary" style="float: right;" size="small" @click="updateCost"> 更新 </el-button>
              <el-button slot="reference" type="primary" icon="el-icon-edit" size="small" circle  />
            </el-popover>
            <el-popconfirm
              title="確定刪除?"
              @onConfirm="deleteCost(scope.row)"
            >
              <el-button slot="reference" style="margin-left: 8px" type="danger" icon="el-icon-delete" size="small" circle />
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
    <CreateCost ref="createCost" :visible.sync="createCostVisible" />
  </div>
</template>

<script>
import CreateCost from './createCost.vue'
import { objClone } from '@/utils/index'
import CostForm from './costForm'

export default {
  components: {
    CreateCost,
    CostForm
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
      createCostVisible: false,
      form: {
        id: 0,
        price: 0,
        count: 0,
        order_date: new Date()
      }
    }
  },
  computed: {
    costList() {
      const list = this.$store.getters['makeup/costList']
      return list
    }
  },
  async created() {
  },
  methods: {
    async deleteCost(cost) {
      try {
        this.loading = true
        await this.$store.dispatch('makeup/deleteMakeupCost', cost.id)
        this.$message.success('刪除成功')
        this.loading = false
      } catch (error) {
        this.$message.error(error)
        this.loading = false
      }
    },
    createCost() {
      this.createCostVisible = true
      this.$refs['createCost'].setDefault()
    },
    setUpdateData(data) {
      this.form = objClone(data)
    },
    async updateCost() {
      const valid = await this.$refs['form'].validate()
      if (valid) {
        try {
          const data = objClone(this.form)
          data.order_date = this.moment(data.order_date).format('YYYY-MM-DD')
          await this.$store.dispatch('makeup/updateMakeupCost', data)
          this.$message.success('更新成功')
        } catch (error) {
          this.$message.error(error)
        }
      }
    },
    close() {
      this.$emit('update:visible', false)
    }
  }
}
</script>

<style lang="scss" scoped>
  .tooltips {
    float: right;
  }
</style>
