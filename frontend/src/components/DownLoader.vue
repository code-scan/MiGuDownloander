<template>
  <el-container>
    <el-header>
      <el-input
        placeholder="请输入内容"
        @keyup.enter.native="keyEnter"
        v-model="keyword"
      >
        <i slot="prefix" class="el-input__icon el-icon-search"></i>
      </el-input>
    </el-header>
    <el-main>
      <div class="tables">
        <el-table
          ref="multipleTable"
          :data="tableData"
          tooltip-effect="dark"
          style="width: 100%"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55"> </el-table-column>

          <el-table-column prop="Name" label="歌曲名称"> </el-table-column>
          <el-table-column prop="Album" label="专辑"> </el-table-column>
          <el-table-column
            prop="Singer"
            label="歌手"
            width="120"
            show-overflow-tooltip
          >
          </el-table-column>
        </el-table>
      </div>
    </el-main>
    <el-footer>
      <el-row>
        <el-col :span="4">
          <div style="margin: 0px">
            <el-button
              :disabled="disbale"
              @click="StartDownload()"
              type="success"
              >开始下载</el-button
            >
          </div>
        </el-col>
        <el-col :span="20">
          <el-progress
            style="margin-top: 15px"
            :text-inside="true"
            :stroke-width="26"
            :percentage="percentage"
          ></el-progress
        ></el-col>
      </el-row>
    </el-footer>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      disbale: false,
      signal: "",
      percentage: 0,
      progressSignal: "",
      tableData: [],
      multipleSelection: [],
      keyword: "",
      message: " ",
    };
  },
  methods: {
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    inertToTable(table) {
      var that = this;
      console.log(table);
      table.forEach((element) => {
        var exists = false;
        that.tableData.forEach((db) => {
          if (db.ID == element.ID) {
            exists = true;
            return;
          }
        });
        if (exists == false) {
          that.tableData.push(element);
        }
      });
    },
    keyEnter() {
      var that = this;
      this.tableData=[]
      this.signal = setInterval(() => {
        window.backend.MiGu.GetResult().then((resp) => {
          that.inertToTable(resp);
        });
        window.backend.MiGu.GetProgress().then((resp) => {
          that.percentage = resp * 100;
        });
      }, 2000);
      window.backend.MiGu.Search(this.keyword).then((resp) => {
        that.inertToTable(resp);
        that.percentage = 100;
        clearInterval(that.signal);
      });

      // start search
    },
    StartDownload() {
      var that = this;
      that.disbale = true;
      that.percentage = 0;
      this.signal = setInterval(() => {
        window.backend.MiGu.GetProgress().then((resp) => {
          console.log(resp);
          if (resp == false) {
            return;
          }
          that.percentage = resp * 100;
        });
      }, 2000);
      window.backend.MiGu.BatchDownload(
        JSON.stringify(this.multipleSelection)
      ).then(() => {
        that.percentage = 100;
        that.disbale = false;
        clearInterval(that.signal);
      });

      console.log(this.multipleSelection);
      // StartDownload
    },
    getMessage: function () {
      var self = this;
      window.backend.basic().then((result) => {
        self.message = result;
      });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-header,
.el-footer {
  background-color: #b3c0d1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: #d3dce6;
  color: #333;
  text-align: center;
  line-height: 200px;
}

.el-main {
  background-color: #e9eef3;
  color: #333;
  height: 100%;
  /* text-align: center;
  line-height: 160px; */
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
.tables {
  height: 500px;
  max-height: 500px;
}

</style>
