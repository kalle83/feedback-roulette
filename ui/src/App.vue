<template>
  <v-app>
    <v-app-bar color="deep-purple" dark app>
      <v-app-bar-title>Feedback Carousel</v-app-bar-title>
      <v-spacer/>
    </v-app-bar>
    <v-main class="pt-20">
      <Question />
    </v-main>
  </v-app>
</template>

<script>
import Question from "./components/QuestionText";

export default {
  name: "App",

  components: {
    Question,
  },
};
</script>
