<template>
  <div>
    <el-dialog title="銷售列表" v-if="visible" :visible="true" @close="close">
      <div class="tooltips">
        <el-button type="primary" size="small" @click="createSale">新增銷售</el-button>
      </div>
      <el-table
        v-loading="loading"
        :data="saleList"
        border
        fit
        highlight-current-row
      >
        <el-table-column align="center" label="銷售" width="160">
          <template slot-scope="scope">
            {{ $_toCurrency(scope.row.price) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="數量" width="80">
          <template slot-scope="scope">
            {{ scope.row.count }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="賣出日期">
          <template slot-scope="scope">
            {{ scope.row.sold_date }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作">
          <template slot-scope="scope">
            <!-- <el-button type="primary" icon="el-icon-edit" size="small" circle @click="updateSale(scope.row)" /> -->
            <el-popover
              placement="right"
              width="400"
              trigger="click"
              @show="setUpdateData(scope.row)"
              >
              <SaleForm ref="form" :form.sync="form" />
              <el-button type="primary" style="float: right;" size="small" @click="updateSale"> 更新 </el-button>
              <el-button slot="reference" type="primary" icon="el-icon-edit" size="small" circle  />
            </el-popover>
            <el-popconfirm
              title="確定刪除?"
              @onConfirm="deleteSale(scope.row)"
            >
              <el-button slot="reference" style="margin-left: 8px" type="danger" icon="el-icon-delete" size="small" circle />
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    <!-- <div slot="footer" class="dialog-footer">
      <el-button @click="close">取消</el-button>
    </div> -->
    </el-dialog>
    <CreateSale ref="createSale" :visible.sync="createSaleVisible" />
  </div>
</template>

<script>
import CreateSale from './createSale'
import { objClone } from '@/utils/index'
import SaleForm from './saleForm'

export default {
  components: {
    CreateSale,
    SaleForm
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
      createSaleVisible: false,
      form: {
        makeup_id: 0,
        price: 0,
        count: 0,
        sold_date: new Date()
      }
    }
  },
  computed: {
    saleList() {
      const list = this.$store.getters['makeup/saleList']
      return list
    }
  },
  async created() {
  },
  methods: {
    async deleteSale(sale) {
      try {
        this.loading = true
        await this.$store.dispatch('makeup/deleteMakeupSale', sale.id)
        this.$message.success('刪除成功')
        this.loading = false
      } catch (error) {
        this.$message.error(error)
        this.loading = false
      }
    },
    createSale() {
      this.createSaleVisible = true
      this.$refs['createSale'].setDefault()
    },
    setUpdateData(data) {
      this.form = objClone(data)
    },
    async updateSale() {
      const valid = await this.$refs['form'].validate()
      if (valid) {
        try {
          const data = objClone(this.form)
          data.sold_date = this.moment(data.sold_date).format('YYYY-MM-DD')
          await this.$store.dispatch('makeup/updateMakeupSale', data)
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
