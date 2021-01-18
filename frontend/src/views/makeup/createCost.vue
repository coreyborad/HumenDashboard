<template>
  <el-dialog width="30%" title="成本列表" v-if="visible" :visible="true" :close-on-click-modal="false" @close="close">
    <CostForm ref="form" :form.sync="form" />
    <div slot="footer" class="dialog-footer">
      <el-button type="primary" @click="save()">建立</el-button>
      <el-button @click="close">取消</el-button>
    </div>
  </el-dialog>
</template>

<script>
import { objClone } from '@/utils/index'
import CostForm from './costForm'

export default {
  components: {
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
      form: {
        makeup_id: 0,
        price: 0,
        count: 0,
        order_date: new Date()
      }
    }
  },
  computed: {
  },
  async created() {
  },
  methods: {
    setDefault() {
      this.form = {
        makeup_id: 0,
        price: 0,
        count: 0,
        order_date: new Date()
      }
    },
    async save() {
      const valid = await this.$refs['form'].validate()
      if (valid) {
        try {
          const data = objClone(this.form)
          data.makeup_id = this.$store.state.makeup.id
          data.order_date = this.moment(data.order_date).format('YYYY-MM-DD')
          await this.$store.dispatch('makeup/createMakeupCost', data)
          this.$message.success('新增成功')
          this.$emit('update:visible', false)
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
