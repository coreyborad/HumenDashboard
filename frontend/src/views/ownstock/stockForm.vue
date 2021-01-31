<template>
    <el-form ref="form" :model="form" label-position="top" :rules="rules" label-width="100px">
      <el-form-item label="股票" prop="stock_number">
        <div
          style="width: 100%;"
        >
          <v-select
            label="stock_name"
            :clearable="false"
            :options="filterList"
            :filterable="false"
            :value="form.stock_number"
            @input="setSelected"
            @search="query => search = query"
          >
            <template #option="{ stock_name, stock_number }">
              {{ stock_name + '(' + stock_number + ')' }}
            </template>
          </v-select>
        </div>
      </el-form-item>
      <el-form-item label="持有股數" prop="shares">
        <el-input-number
          style="width: 100%;"
          v-model="form.shares"
          controls-position="right"
          :min="1"
        />
      </el-form-item>
      <el-form-item label="持有總成本" prop="cost">
        <el-input-number
          style="width: 100%;"
          v-model="form.cost"
          :precision="2"
          controls-position="right"
          :min="0"
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
          stock_number: null,
          shares: 0,
          cost: 0
        }
      }
    }
  },
  computed: {
    filterList() {
      return this.$store.state.stock.stock_list.filter(stock => {
        if (this.search === '') {
          return true
        }
        if (stock.stock_number.indexOf(this.search) !== -1) {
          return true
        }
        if (stock.stock_name.indexOf(this.search) !== -1) {
          return true
        }
        return false
      })
    }
  },
  data() {
    return {
      search: '',
      rules: {
        stock_number: [
          { required: true, message: '請選擇股票', trigger: 'blur' }
        ],
        shares: [
          { required: true, message: '請輸入股數', trigger: 'blur' }
        ],
        cost: [
          { required: true, message: '請輸入成本', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    setSelected(value) {
      this.form.stock_number = value.stock_number
    },
    validate() {
      return this.$refs['form'].validate()
    }
  }
}
</script>

<style>

</style>