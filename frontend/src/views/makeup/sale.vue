<template>
  <div>
    <el-dialog title="銷售列表" :visible="visible" @close="close">
      <div class="tooltips">
        <el-button type="primary" size="small" @click="createSaleVisible = true">新增銷售</el-button>
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
            <el-button type="primary" icon="el-icon-edit" size="small" circle @click="updateSale(scope.row)" />
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
    <CreateSale :visible.sync="createSaleVisible" />
  </div>
</template>

<script>
import CreateSale from './createSale'

export default {
  components: {
    CreateSale
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
      createSaleVisible: false
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
    updateSale(sale) {
      console.log(sale)
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
