<template>
  <el-dialog title="新增持有股票" :visible="visible" @close="close">
    <el-form :model="form">
      <el-form-item label="股票" :label-width="formLabelWidth" required>
        <div
          style="width: 180px;
                line-height: 38px;"
        >
          <v-select
            label="stock_name"
            :clearable="false"
            :options="filterList"
            :filterable="false"
            @input="setSelected"
            @search="query => search = query"
          >
            <template #option="{ stock_name, stock_number }">
              {{ stock_name + '(' + stock_number + ')' }}
            </template>
          </v-select>
        </div>

      </el-form-item>
      <el-form-item label="持有股數" :label-width="formLabelWidth" required>
        <el-input-number
          v-model="form.shares"
          controls-position="right"
          :min="1"
        />
      </el-form-item>
      <el-form-item label="持有總成本" :label-width="formLabelWidth" required>
        <el-input-number
          v-model="form.cost"
          :precision="2"
          controls-position="right"
          :min="0"
        />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button type="primary" @click="add">確定</el-button>
      <el-button @click="close">取消</el-button>
    </div>
  </el-dialog>
</template>

<script>
import { getStockList, createUserStock } from '@/api/stock'

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
      formLabelWidth: '120px',
      stockList: [],
      search: '',
      form: {
        stock_number: null,
        shares: 0,
        cost: 0
      }
    }
  },
  computed: {
    filterList() {
      return this.stockList.filter(stock => {
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
  async created() {
    this.stockList = await getStockList()
    // this.filterList = Object.assign(this.stockList)
  },
  methods: {
    setSelected(value) {
      this.form.stock_number = value.stock_number
    },
    async add() {
      try {
        await createUserStock(this.form)
        this.$emit('add')
        this.close()
      } catch (error) {
        this.$message.error(error)
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
