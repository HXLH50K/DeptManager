<template>
  <div>
    <el-container style="height: 500px; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
        <el-menu router :default-openeds="['0']">
          <el-submenu
            v-for="(item, index) in $router.options.routes"
            v-bind:key="item.id"
            v-bind:index="index + ''"
          >
            <template slot="title">
              <i class="el-icon-message"></i>{{ item.name }}
            </template>
            <el-menu-item
              v-for="item2 in item.children"
              v-bind:key="item2.id"
              v-bind:index="item2.path"
              v-bind:class="$route.path == item2.path ? 'is-active' : ''"
              v-show="item2.name != '部门编辑'"
            >
              {{ item2.name }}
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </div>
</template>
<script>

export default {
  name: "Index",
}
</script>

<style>
</style>