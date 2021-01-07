<template>
  <div>
    <el-dialog title="成本列表" :visible="visible" :close-on-click-modal="false" @close="close">
      <div class="tooltips">
        <el-button type="primary" size="small" @click="createCostVisible = true">新增進貨</el-button>
      </div>
      <el-table
        v-loading="loading"
        :data="costList"
        border
        fit
        highlight-current-row
      >
        <el-table-column align="center" label="成本" width="160">
          <template slot-scope="scope">
            {{ $_toCurrency(scope.row.price) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="數量" width="80">
          <template slot-scope="scope">
            {{ scope.row.count }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="進貨日期">
          <template slot-scope="scope">
            {{ scope.row.order_date }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作">
          <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-edit" size="small" circle @click="updateCost(scope.row)" />
            <el-button type="danger" icon="el-icon-delete" size="small" circle @click="deleteCost(scope.row)" />
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
    <CreateCost :visible.sync="createCostVisible" />
  </div>
</template>

<script>
import CreateCost from './createCost.vue'

export default {
  components: {
    CreateCost
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
      createCostVisible: false
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
    deleteCost(cost) {
      console.log(cost)
    },
    updateCost(cost) {
      console.log(cost)
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
