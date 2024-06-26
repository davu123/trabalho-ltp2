<template>
  <div class="box">
    <div class="rectangle"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Box",
});
</script>

<style>
.box {
  height: 700px;
  width: 1540px;
}

.box .rectangle {
  background-color: #ffffffcc;
  border-radius: 45px;
  height: 700px;
  left: 0;
  position: fixed;
  top: 0;
  width: 1540px;
}
</style>
