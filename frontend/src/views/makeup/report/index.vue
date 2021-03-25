<template>
  <div>
    <el-row :gutter="12">
      <el-col :span="12">
        <el-card
          shadow="always"
          style="text-align:center;"
          v-loading="monthSaleBar.loading"
        >
          <div slot="header" class="clearfix">
            <div class="card-header-title">月份銷售狀況報表</div>
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
        <el-card 
          shadow="always"
          style="text-align:center;"
          v-loading="yearsSalePie.loading"
        >
          <div slot="header" class="clearfix">
            <div class="card-header-title">年度銷售狀況報表</div>
            <el-date-picker
              v-model="yearsSalePie.year"
              type="year"
              placeholder="選擇年份"
              value-format="yyyy">
            </el-date-picker>
            <el-button type="primary" @click="onYearPie">送出</el-button>
          </div>
          <div style="margin:auto;">
            <v-chart :options="yearsSalePie.charts" style="width:100%;" autoresize/>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { getReportsByQuery } from '@/api/makeup'

export default {
  name: 'report',
  data() {
    return {
      monthSaleBar: {
        loading: false,
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
      },
      yearsSalePie: {
        loading: false,
        year: this.moment().format('YYYY'),
        charts: {
          tooltip: {
            trigger: 'item',
          },
          legend: {
          },
          series: [
            {
              name: '銷售額',
              avoidLabelOverlap: false,
              label: {
                  show: false,
                  position: 'center'
              },
              data: [
                {value: 335, name: '直接访问'},
                {value: 310, name: '邮件营销'},
                {value: 274, name: '联盟广告'},
                {value: 235, name: '视频广告'},
                {value: 400, name: '搜索引擎'}
              ],
              type: 'pie',
            },
          ]
        }
      }
    }
  },
  async mounted() {
    this.onMonthSaleBar()
    this.onYearPie()
  },
  methods: {
    async onMonthSaleBar() {
      this.monthSaleBar.loading = true
      const costs = await getReportsByQuery({
        'type': 'a',
        'target': 'cost',
        'date_start': this.monthSaleBar.startMonth,
        'date_end': this.monthSaleBar.endMonth
      })
      const sales = await getReportsByQuery({
        'type': 'a',
        'target': 'sale',
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
      this.monthSaleBar.loading = false
    },
    async onYearPie() {
      this.yearsSalePie.loading = true
      const sales = await getReportsByQuery({
        'type': 'a',
        'target': 'sale',
        'date_start': this.moment(this.yearsSalePie.year).startOf('year').format('YYYY-MM'),
        'date_end': this.moment(this.yearsSalePie.year).endOf('year').format('YYYY-MM'),
      })
      const data = sales.map( s => {
        return {
          name: s.month,
          value: s.price
        }
      })
      this.yearsSalePie.charts.series[0].data = data
      this.yearsSalePie.loading = false
    }
  }
}
</script>

<style lang="scss" scoped>
  .card-header-title {
    margin-bottom: 8px;
  }
</style>
