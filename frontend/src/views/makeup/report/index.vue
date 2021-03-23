<template>
  <div>
    <el-row :gutter="12">
      <el-col :span="12">
        <el-card shadow="always" style="text-align:center;">
          <div slot="header" class="clearfix">
            <el-date-picker
              v-model="monthSaleBar.startMonth"
              type="month"
              placeholder="選擇初始月"
              value-format="yyyy-MM">
            </el-date-picker>
            <el-date-picker
              v-model="monthSaleBar.endMonth"
              type="month"
              placeholder="選擇結束月"
              value-format="yyyy-MM">
            </el-date-picker>
            <el-button type="primary" @click="onMonthSaleBar">送出</el-button>
          </div>
          <div style="margin:auto;">
            <v-chart :options="monthSaleBar.charts" style="width:100%;" autoresize/>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover">
          建置中
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { getMakeupCostMonthByQuery, getMakeupSaleMonthByQuery } from '@/api/makeup'
import 'echarts/lib/component/legend'

export default {
  name: 'report',
  data() {
    return {
      monthSaleBar: {
        startMonth: this.moment().add(-11, 'months').format('YYYY-MM'),
        endMonth: this.moment().format('YYYY-MM'),
        charts: {
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'shadow'
            },
            formatter: datas => {
              return `
              ${datas[0].axisValueLabel}<br/>
              ${datas[0].marker} ${datas[0].seriesName} ${this.$_toCurrency(datas[0].value)}<br/>
              ${datas[1].marker} ${datas[1].seriesName} ${this.$_toCurrency(datas[1].value)}<br/>
              淨額: ${this.$_toCurrency(datas[1].value - datas[0].value)}
              `
            }
          },
          legend: {
          },
          xAxis: {
            type: 'category',
            data: [],
            axisLabel: {
              interval: 0,
              rotate: 15
            }
          },
          yAxis: {
            type: 'value'
          },
          series: [
            {
              name: '成本',
              data: [],
              type: 'bar',
            },
            {
              name: '銷售',
              data: [],
              type: 'bar',
            }
          ]
        }
      }
    }
  },
  async mounted() {
    this.onMonthSaleBar()
  },
  methods: {
    async onMonthSaleBar() {
      const costs = await getMakeupCostMonthByQuery({
        'date_start': this.monthSaleBar.startMonth,
        'date_end': this.monthSaleBar.endMonth
      })
      const sales = await getMakeupSaleMonthByQuery({
        'date_start': this.monthSaleBar.startMonth,
        'date_end': this.monthSaleBar.endMonth
      })
      const xAxis = []
      const costList = []
      
      costs.forEach((element) => {
        xAxis.push(element.month)
        costList.push(element.price)
      });
      this.monthSaleBar.charts.xAxis.data = xAxis
      this.monthSaleBar.charts.series[0].data = costList
      const saleList = []
      sales.forEach((element) => {
        saleList.push(element.price)
      });
      this.monthSaleBar.charts.series[1].data = saleList
    }
  }
}
</script>
