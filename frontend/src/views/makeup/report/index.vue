<template>
  <div style="margin: 0 12px;">
    <el-row :gutter="12" type="flex" class="block" justify="center">
      <el-col :span="12">
        <el-card 
          shadow="always"
          style="text-align:center;"
          v-loading="monthItemTable.loading"
        >
          <div slot="header" class="clearfix">
            <div class="card-header-title">商品月份售出表</div>
            <el-date-picker
              v-model="monthItemTable.month"
              type="month"
              placeholder="選擇月份"
              value-format="yyyy-MM">
            </el-date-picker>
            <el-button type="primary" @click="onMonthItemTable">送出</el-button>
          </div>
          <div style="margin:auto;">
            <el-table
              :data="monthItemTable.data"
              border
              style="width: 100%"
              height="400"
            >
              <el-table-column
                label="品牌"
                width="180">
                <template slot-scope="scope">
                  {{ scope.row.info.brand }}
                </template>
              </el-table-column>
              <el-table-column
                label="名稱"
                width="180">
                <template slot-scope="scope">
                  {{ scope.row.info.name }}
                </template>
              </el-table-column>
              <el-table-column
                label="色號"
                width="180">
                <template slot-scope="scope">
                  {{ scope.row.info.color_name }}
                </template>
              </el-table-column>
              <el-table-column
                prop="count"
                label="賣出數量"
                width="180">
              </el-table-column>
            </el-table>
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
    <el-row :gutter="12" type="flex" class="block" justify="center">
      <el-col :span="24">
        <el-card
          shadow="always"
          style="text-align:center;"
          v-loading="monthSaleBar.loading"
        >
          <div slot="header" class="clearfix">
            <div class="card-header-title">月份成本收入報表</div>
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
    </el-row>
    <el-row :gutter="12" type="flex" class="block" justify="center">
      <el-col :span="24">
        <el-card 
          shadow="always"
          style="text-align:center;"
          v-loading="monthItemBar.loading"
        >
          <div slot="header" class="clearfix">
            <div class="card-header-title">商品月份淨利表</div>
            <el-date-picker
              v-model="monthItemBar.startMonth"
              type="month"
              placeholder="選擇初始月"
              value-format="yyyy-MM">
            </el-date-picker>
            <el-date-picker
              v-model="monthItemBar.endMonth"
              type="month"
              placeholder="選擇結束月"
              value-format="yyyy-MM">
            </el-date-picker>
            <el-button type="primary" @click="onMonthItemBar">送出</el-button>
          </div>
          <div style="margin:auto;">
            <v-chart :options="monthItemBar.charts" style="width:100%;" autoresize/>
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
              淨支出: ${this.$_toCurrency(datas[1].value - datas[0].value)}
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
              name: '支出',
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
              ],
              type: 'pie',
            },
          ]
        }
      },
      monthItemBar: {
        loading: false,
        startMonth: this.moment().add(-11, 'months').format('YYYY-MM'),
        endMonth: this.moment().format('YYYY-MM'),
        charts: {
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'shadow'
            }
          },
          xAxis: {
            type: 'category',
            data: [],
            axisLine: {onZero: true},
          },
          grid: {
            containLabel: true
          },
          yAxis: [
              {
                type: 'value',
                name: '金額',
                min: -50000,
                max: 50000,
                interval: 10000,
                axisLabel: {
                  formatter: '${value}'
                }
              },
              {
                type: 'value',
                name: '毛利率',
                min: -60,
                max: 60,
                interval: 12,
                axisLabel: {
                  formatter: '{value}%'
                }
              }
          ],
          legend: {
            data: ['營收', '成本', '淨利'],
          },
          series: [
            {
              name: '營收',
              type: 'bar',
              stack: 'one',
              emphasis: {
                focus: 'series'
              },
              label: {
                show: true,
                position: 'top',
                formatter: item => {
                  return `${this.$_toCurrency(item.value)}`
                }
              },
              data: []
            },
            {
              name: '成本',
              type: 'bar',
              stack: 'one',
              emphasis: {
                focus: 'series'
              },
              label: {
                show: true,
                position: 'bottom',
                formatter: item => {
                  return `${this.$_toCurrency(item.value)}`
                }
              },
              data: []
            },
            {
              name: '淨利',
              type: 'bar',
              stack: 'two',
              emphasis: {
                focus: 'series'
              },
              label: {
                show: true,
                position: 'bottom',
                formatter: item => {
                  return `${this.$_toCurrency(item.value)}`
                }
              },
              data: []
            },
            {
              name: '毛利率',
              type: 'line',
              label: {
                show: true,
                position: 'top',
                formatter: item => {
                  return `${item.value}%`
                }
              },
              yAxisIndex: 1,
              data: []
            },
          ]
        }
      },
      monthItemTable:{
        loading: false,
        month: this.moment().format('YYYY-MM'),
        data: []
      }
    }
  },
  async mounted() {
    this.onMonthSaleBar()
    this.onYearPie()
    this.onMonthItemBar()
    this.onMonthItemTable()
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
      const data = sales.map(s => {
        return {
          name: s.month,
          value: s.price
        }
      })
      this.yearsSalePie.charts.series[0].data = data
      this.yearsSalePie.loading = false
    },
    async onMonthItemBar() {
      this.monthItemBar.loading = true
      const list = await getReportsByQuery({
        'type': 'c',
        'date_start': this.monthItemBar.startMonth,
        'date_end': this.monthItemBar.endMonth
      })
      this.monthItemBar.charts.xAxis.data = []
      this.monthItemBar.charts.series[0].data = []
      this.monthItemBar.charts.series[1].data = []
      this.monthItemBar.charts.series[2].data = []
      this.monthItemBar.charts.series[3].data = []
      list.forEach(l => {
        this.monthItemBar.charts.xAxis.data.push(l.month)
        this.monthItemBar.charts.series[0].data.push(l.sale)
        this.monthItemBar.charts.series[1].data.push(l.cost * -1)
        this.monthItemBar.charts.series[2].data.push(l.sale - l.cost)
        if(l.sale <= 0 || l.sale < l.cost ){
          this.monthItemBar.charts.series[3].data.push(0)
        }else{
          this.monthItemBar.charts.series[3].data.push(Math.round(((l.sale - l.cost) / l.sale) * 100))
        }
      })
      this.monthItemBar.loading = false
    },
    async onMonthItemTable() {
      this.monthItemTable.loading = true
      this.monthItemTable.data = await getReportsByQuery({
        'type': 'd',
        'date': this.monthItemTable.month
      })
      this.monthItemTable.loading = false
    }
  }
}
</script>

<style lang="scss" scoped>
  .card-header-title {
    margin-bottom: 8px;
  }
  .block {
    margin: 12px 0px;
  }
</style>
