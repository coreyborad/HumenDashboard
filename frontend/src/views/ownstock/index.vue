<template>
  <div class="app-container">
    <div class="tooltips">
      <el-button type="primary" size="small" @click="stockDialogVisible = true">新增持有股票</el-button>
    </div>
    <el-table
      v-loading="loading"
      :data="list"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="股票資訊">
        <template slot-scope="scope">
          {{ scope.row.stock_info.stock_name + '(' + scope.row.stock_info.stock_number + ')' }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="持有股數">
        <template slot-scope="scope">
          {{ scope.row.shares }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="持有總成本">
        <template slot-scope="scope">
          {{ $_toCurrency(scope.row.cost) }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="目前損益">
        <template slot-scope="scope">
          {{ $_toCurrency((scope.row.last_stock.price_on_close * scope.row.shares) - scope.row.cost) }}
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
            <StockForm ref="form" :form.sync="form" />
            <el-button type="primary" style="float: right;" size="small" @click="updateStock"> 更新 </el-button>
            <el-button slot="reference" type="primary" icon="el-icon-edit" size="small" circle  />
          </el-popover>
          <el-button type="danger" style="margin-left: 8px" icon="el-icon-delete" size="small" circle @click="deleteItem(scope.row.id)" />
        </template>
      </el-table-column>
    </el-table>
    <CreateDialog :visible.sync="stockDialogVisible"/>
  </div>
</template>

<script>
import { objClone } from '@/utils/index'
import StockForm from './stockForm'
import CreateDialog from './create.vue'

export default {
  components: {
    CreateDialog,
    StockForm
  },
  data() {
    return {
      loading: false,
      stockDialogVisible: false,
      form: {
        id: 0,
        stock_number: null,
        shares: 0,
        cost: 0
      }
    }
  },
  computed: {
    list() {
      return this.$store.state.stock.list
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.loading = true
      await this.$store.dispatch('stock/getList')
      await this.$store.dispatch('stock/getStockList')
      this.loading = false
    },
    setUpdateData(data) {
      this.form.id = data.id
      this.form.stock_number = data.stock_number
      this.form.shares = data.shares
      this.form.cost = data.cost
    },
    async updateStock() {
      const valid = await this.$refs['form'].validate()
      if (valid) {
        try {
          await this.$store.dispatch('stock/updateUserStock', this.form)
          this.$message.success('更新成功')
        } catch (error) {
          this.$message.error(error)
        }
      }
    },
    async deleteItem(id) {
      this.loading = true
      try {
        await this.$store.dispatch('stock/deleteUserStockInfo', id)
      } catch (error) {
        this.$message.error(error)
      }
      this.loading = false
    }
  }
}
</script>

<style lang="scss" scoped>
  .tooltips {
    margin-bottom: 16px;
    float:right;
  }
</style>
