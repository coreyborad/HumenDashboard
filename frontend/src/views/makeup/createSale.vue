<template>
  <el-dialog width="30%" title="銷售列表" v-if="visible" :visible="true" :close-on-click-modal="false" @close="close">
    <el-form ref="form" :model="form" label-position="top" :rules="rules" label-width="100px">
      <el-form-item label="銷售單價" prop="price">
        <el-input-number
          v-model="form.price"
          style="width: 100%;"
          controls-position="right"
          :min="0"
        />
      </el-form-item>
      <el-form-item label="數量" prop="count">
        <el-input-number
          v-model="form.count"
          style="width: 100%;"
          controls-position="right"
          :min="0"
        />
      </el-form-item>
      <el-form-item label="銷售日期" prop="sold_date">
        <el-date-picker
          v-model="form.sold_date"
          style="width: 100%;"
          type="date"
          placeholder="選擇日期"
        />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button type="primary" @click="save()">建立</el-button>
      <el-button @click="close">取消</el-button>
    </div>
  </el-dialog>
</template>

<script>
import { objClone } from '@/utils/index'

export default {
  components: {
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
      rules: {
        price: [
          { required: true, message: '請輸入文字', trigger: 'blur' }
        ],
        count: [
          { required: true, message: '請輸入文字', trigger: 'blur' }
        ],
        sold_date: [
          { required: true, message: '請輸入文字', trigger: 'blur' }
        ]
      },
      form: {
        makeup_id: 0,
        price: 0,
        count: 0,
        sold_date: new Date()
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
    setDefault() {
      this.form = {
        makeup_id: 0,
        price: 0,
        count: 0,
        sold_date: new Date()
      }
    },
    async save() {
      const valid = await this.$refs['form'].validate()
      if (valid) {
        try {
          const data = objClone(this.form)
          data.makeup_id = this.$store.state.makeup.id
          data.sold_date = this.moment(data.sold_date).format('YYYY-MM-DD')
          await this.$store.dispatch('makeup/createMakeupSale', data)
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
