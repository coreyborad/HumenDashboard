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
          {{ scope.row.cost }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="目前損益">
        <template slot-scope="scope">
          {{ (scope.row.last_stock.price_on_close * scope.row.shares) - scope.row.cost }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button type="danger" icon="el-icon-delete" size="small" circle @click="deleteItem(scope.row.id)" />
        </template>
      </el-table-column>
    </el-table>
    <CreateDialog :visible.sync="stockDialogVisible" @add="fetchData" />
  </div>
</template>

<script>
import { getUserStock, deleteUserStock } from '@/api/stock'

import CreateDialog from './create.vue'

export default {
  components: {
    CreateDialog
  },
  data() {
    return {
      list: null,
      loading: true,
      stockDialogVisible: false
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.loading = true
      this.list = await getUserStock()
      this.loading = false
    },
    async deleteItem(id) {
      this.loading = true
      try {
        await deleteUserStock(id)
        this.fetchData()
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
