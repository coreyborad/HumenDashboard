<template>
  <div class="app-container">
    <div class="tooltips">
      <el-button type="primary" size="small" @click="createItemDialog">新增品項</el-button>
    </div>
    <el-table
      v-loading="loading"
      :data="list"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="品牌">
        <template slot-scope="scope">
          {{ scope.row.brand }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="品項">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="總成本">
        <template slot-scope="scope">
          {{ $_toCurrency(scope.row.cost_total) }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="平均成本">
        <template slot-scope="scope">
          {{ $_toCurrency(scope.row.cost_total / scope.row.cost_count) }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="總銷售額">
        <template slot-scope="scope">
          {{ $_toCurrency(scope.row.sale_total) }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="平均銷售額">
        <template slot-scope="scope">
          {{ $_toCurrency(scope.row.sale_total / scope.row.sale_count) }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="淨利">
        <template slot-scope="scope">
          {{ $_toCurrency(scope.row.sale_total - scope.row.cost_total) }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button type="primary" icon="el-icon-info" size="small" circle @click="openColorDialog(scope.row)" />
        </template>
      </el-table-column>
    </el-table>
    <ColorDialog ref="colorDialog" :visible.sync="colorDialogVisible" />
    <ItemDialog ref="itemDialog" :visible.sync="itemDialogVisible" />
  </div>
</template>

<script>
// import { getMakeup } from '@/api/makeup'

import ColorDialog from './color.vue'
import ItemDialog from './itemDialog.vue'

export default {
  components: {
    ColorDialog,
    ItemDialog
  },
  data() {
    return {
      colorDialogVisible: false,
      itemDialogVisible: false,
      loading: true
    }
  },
  computed: {
    list() {
      return this.$store.state.makeup.list
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.loading = true
      await this.$store.dispatch('makeup/getList')
      this.loading = false
    },
    async openColorDialog(data) {
      await this.$store.dispatch('makeup/setCurrentValue', { target: 'brand', value: data.brand })
      await this.$store.dispatch('makeup/setCurrentValue', { target: 'name', value: data.name })
      this.colorDialogVisible = true
    },
    async createItemDialog() {
      this.itemDialogVisible = true
      this.$refs['itemDialog'].setDefault()
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
