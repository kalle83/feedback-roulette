<template>
  <v-app>
    <Navigation />
    <Content class="mt-4" />
    <Footer />
  </v-app>
</template>

<script>
import Navigation from "./components/Navigation.vue";
import Footer from "./components/Footer.vue";
import Content from "./components/Content.vue";

export default {
  name: "App",

  components: {
    Navigation,
    Content,
    Footer,
  },
};
</script>

<style scoped>
#app {
  background: url("@/assets/background.jpg");
  background-size: cover;
}
</style>
