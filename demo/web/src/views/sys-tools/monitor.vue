<template>
  <div v-loading="loading">
    <BasicLayout>
      <template #wrapper>
        <el-row :gutter="10" class="mb10">
          <el-col :sm="24" :md="8">
            <el-card v-if="info.cpu" class="box-card">
              <div slot="header" class="clearfix">
                <span>CPU使用率</span>
              </div>
              <div class="monitor">
                <div class="monitor-header">
                  <el-progress :color="$store.state.settings.theme" type="circle" :percentage="info.cpu.Percent" />
                </div>
                <div class="monitor-footer">
                  <Cell label="CPU主频" :value="info.cpu.cpuInfo[0].modelName.split('@ ')[1]" border />
                  <Cell label="核心数" :value="`${info.cpu.cpuInfo[0].cores}`" />
                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :sm="24" :md="8">
            <el-card v-if="info.mem" class="box-card">
              <div slot="header" class="clearfix">
                <span>内存使用率</span>
              </div>
              <div class="monitor">
                <div class="monitor-header">
                  <el-progress :color="$store.state.settings.theme" type="circle" :percentage="info.mem.usage" />
                </div>
                <div class="monitor-footer">
                  <Cell label="总内存" :value="info.mem.total+'G'" border />
                  <Cell label="已用内存" :value="info.mem.used+'G'" />
                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :sm="24" :md="8">
            <el-card v-if="info.disk" class="box-card">
              <div slot="header" class="clearfix">
                <span>磁盘信息</span>
              </div>
              <div class="monitor">
                <div class="monitor-header">
                  <el-progress :color="$store.state.settings.theme" type="circle" :percentage=" Number(( (info.disk.total-info.disk.free) / info.disk.total * 100).toFixed(2))" />
                </div>
                <div class="monitor-footer">
                  <Cell label="总磁盘" :value="info.disk.total+'G'" border />
                  <Cell label="已用磁盘" :value="info.disk.total-info.disk.free+'G'" />
                </div>
              </div>
            </el-card>
          </el-col>
        </el-row>

        <el-card v-if="info.os" class="box-card">
          <div slot="header" class="clearfix">
            <span>go运行环境</span>
          </div>
          <div class="monitor">
            <Cell label="GO 版本" :value="info.os.version" border />
            <Cell label="Goroutine" :value="`${info.os.numGoroutine}`" border />
            <Cell label="项目地址" :value="info.os.projectDir" />
          </div>
        </el-card>

        <el-card v-if="info.os" class="box-card">
          <div slot="header" class="clearfix">
            <span>服务器信息</span>
          </div>
          <div class="monitor">
            <Cell label="主机名称" :value="info.os.hostName" border />
            <Cell label="操作系统" :value="info.os.goOs" border />
            <Cell label="服务器IP" :value="info.os.ip" border />
            <Cell label="系统架构" :value="info.os.arch" border />
            <Cell label="CPU" :value="info.cpu.cpuInfo[0].modelName" border />
            <Cell label="当前时间" :value="info.os.time" />
          </div>
        </el-card>

        <el-card>
          <div slot="header">
            <span>磁盘状态</span>
          </div>
          <el-table :data="info.diskList">
            <el-table-column label="盘符路径" prop="path" />
            <el-table-column label="文件系统" prop="fstype" />
            <el-table-column label="总大小" prop="total" />
            <el-table-column label="可用大小" prop="free" />
            <el-table-column label="已用大小" prop="used" />
            <el-table-column label="已用百分比">
              <template slot-scope="scope">
<!--                <div class="cell" :class="{'text-danger': scope.row.usedPercent > 80}">{{ scope.row.usedPercent }}%</div>-->
                <el-tag type="success" v-if="scope.row.usedPercent < 80">{{ scope.row.usedPercent }}%</el-tag>
                <el-tag type="danger" v-else>{{ scope.row.usedPercent }}%</el-tag>
              </template>
            </el-table-column>
          </el-table>
<!--          <div class="el-table el-table&#45;&#45;enable-row-hover el-table&#45;&#45;medium">-->
<!--            <table cellspacing="0" style="width: 100%;">-->
<!--              <thead>-->
<!--                <tr>-->
<!--                  <th class="is-leaf"><div class="cell">盘符路径</div></th>-->
<!--                  <th class="is-leaf"><div class="cell">文件系统</div></th>-->
<!--                  <th class="is-leaf"><div class="cell">总大小</div></th>-->

<!--                  <th class="is-leaf"><div class="cell">可用大小</div></th>-->
<!--                  <th class="is-leaf"><div class="cell">已用大小</div></th>-->
<!--                  <th class="is-leaf"><div class="cell">已用百分比</div></th>-->
<!--                </tr>-->
<!--              </thead>-->
<!--              <tbody v-if="info.diskList">-->
<!--                <tr v-for="(forList,index) in info.diskList" :key="index">-->
<!--                  <td><div class="cell">{{ forList.path }}</div></td>-->
<!--                  <td><div class="cell">{{ forList.fstype }}</div></td>-->
<!--                  <td><div class="cell">{{ forList.total }}M</div></td>-->
<!--                  <td><div class="cell">{{ forList.free }}M</div></td>-->
<!--                  <td><div class="cell">{{ forList.used }}M</div></td>-->
<!--                  <td><div class="cell" :class="{'text-danger': forList.usedPercent > 80}">{{ forList.usedPercent }}%</div></td>-->
<!--                </tr>-->
<!--              </tbody>-->
<!--            </table>-->
<!--          </div>-->
        </el-card>

      </template>
    </BasicLayout>
  </div>
</template>

<script>
import Cell from '@/components/Cell/index'
import { getServer } from '@/api/monitor/server'
export default {
  name: 'Monitor',
  components: {
    Cell
  },
  data() {
    return {
      loading: true,
      info: {}
    }
  },
  created() {
    this.getServerInfo()
  },
  methods: {
    getServerInfo() {
      getServer().then(ret => {
        if (ret.code === 200) {
          this.loading = false
          this.info = ret
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
  .monitor{
    .monitor-header{
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }
  .el-card__body{
    padding: 20px 20px 0 20px!important;
  }
</style>
