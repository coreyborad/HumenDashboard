<template>
  <el-form ref="form" :model="form" label-position="top" :rules="rules" label-width="100px">
    <el-form-item label="單位成本" prop="price">
      <el-input-number
        v-model="form.price"
        style="width: 100%;"
        controls-position="right"
        :min="0"
      />
    </el-form-item>
    <el-form-item label="數量" prop="count">
      <el-input-number
        :disabled="form.update"
        v-model="form.count"
        style="width: 100%;"
        controls-position="right"
        :min="0"
      />
    </el-form-item>
    <el-form-item label="進貨日期" prop="order_date">
      <el-date-picker
        v-model="form.order_date"
        style="width: 100%;"
        type="date"
        placeholder="選擇日期"
      />
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  props: {
    form: {
      type: Object,
      default: () =>{
        return {
          makeup_id: 0,
          price: 0,
          count: 0,
          order_date: new Date()
        }
      }
    }
  },
  data() {
    return {
      rules: {
        price: [
          { required: true, message: '請輸入文字', trigger: 'blur' }
        ],
        count: [
          { required: true, message: '請輸入文字', trigger: 'blur' }
        ],
        order_date: [
          { required: true, message: '請輸入文字', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    validate() {
      return this.$refs['form'].validate()
    }
  }
}
</script>

<style>

</style>