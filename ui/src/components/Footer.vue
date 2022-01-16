<template>
  <v-footer padless color="#132577" dark>
    <v-col class="text-center" cols="12">
      {{ new Date().getFullYear() }} — <strong>{{ stateData }}</strong>
    </v-col>
  </v-footer>
</template>

<script>
export default {
  data() {
    return {
      stateData: "",
    };
  },
  mounted() {
    this.stateData = this.$store.getters.getFooterMsg;
  },
};
</script>

<style>
</style>