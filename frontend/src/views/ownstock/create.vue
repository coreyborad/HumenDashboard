<template>
  <el-dialog title="新增持有股票" v-if="visible" :visible="true" @close="close">
    <StockForm ref="form" :form.sync="form" />
    <div slot="footer" class="dialog-footer">
      <el-button type="primary" @click="add">儲存</el-button>
      <el-button @click="close">取消</el-button>
    </div>
  </el-dialog>
</template>

<script>
import StockForm from './stockForm'

export default {
  components: {
    StockForm
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
      form: {
        stock_number: null,
        shares: 0,
        cost: 0
      }
    }
  },
  computed: {
  },
  async created() {
  },
  methods: {
    async add() {
      const valid = await this.$refs['form'].validate()
      if (valid) {
        try {
          await this.$store.dispatch('stock/createUserStockInfo', this.form)
          this.$message.success('新增成功')
          this.close()
        } catch (error) {
          this.$message.error(error)
        }
      }
    },
    close() {
      this.form = {
        stock_number: null,
        shares: 0,
        cost: 0
      }
      this.$emit('update:visible', false)
    }
  }
}
</script>
