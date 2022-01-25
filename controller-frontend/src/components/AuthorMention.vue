<template>
  <div class="author-mention">cubid.cloud &copy; {{ new Date().getFullYear() }} <span v-if="!$vuetify.breakpoint.mobile">- made by Jasper S.</span></div>

</template>

<style scoped>
div.author-mention {
  position: fixed;
  bottom: 30px;
  opacity: .25;
  user-select: none;

  left: 50%;
  transform: translateX(-50%);
}
</style>

