<template>
  <el-dialog title="銷售列表" :visible="visible" @close="close">
    <!-- <div class="tooltips">
      <el-button type="primary" size="small" @click="stockDialogVisible = true">新增持有股票</el-button>
    </div> -->
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
      <el-table-column align="center" label="對應的進貨日">
        <template slot-scope="scope">
          {{ scope.row.cost.order_date }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button type="primary" icon="el-icon-edit" size="small" circle @click="updateSale(scope.row)" />
          <el-button type="danger" icon="el-icon-delete" size="small" circle @click="deleteSale(scope.row)" />
        </template>
      </el-table-column>
    </el-table>
    <!-- <div slot="footer" class="dialog-footer">
      <el-button @click="close">取消</el-button>
    </div> -->
  </el-dialog>
</template>

<script>

export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      loading: false,
      saleList: []
    }
  },
  computed: {
  },
  async created() {
  },
  methods: {
    setValue(sale) {
      this.saleList = sale
    },
    deleteSale(sale) {
      console.log(sale)
    },
    updateSale(sale) {
      console.log(sale)
    },
    close() {
      this.$emit('update:visible', false)
    }
  }
}
</script>
