<template>
  <v-app>
    
    <v-main class="pt-8">
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
