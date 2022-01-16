<template>
  <v-container>
    <v-row justify="center">
      <v-col >
        <v-carousel v-model="model">
          <v-carousel-item v-for="question in questionList" :key="question.id">
            <v-sheet :color="question.color" height="100%" tile>
              <v-row class="fill-height" align="center" justify="center">
                <div class="text-h2 white--text text-center px-12">{{ question.text }}</div>
              </v-row>
            </v-sheet>
          </v-carousel-item>
        </v-carousel>
      </v-col>
    </v-row>
    <v-row class="pt-9" justify="center">
      <v-col cols="6">
        <v-slider
          v-model="members"
          class="align-self-stretch"
          min="3"
          max="10"
          ticks="always"
          tick-size="4"
          step="1"
          thumb-label="always"
        >
          <template v-slot:append>
            <v-text-field
              v-model="members"
              class="mt-0 pt-0"
              hide-details
              single-line
              type="number"
              style="width: 60px"
            ></v-text-field> </template
        ></v-slider>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  name: "QuestionText",
  computed: {
    questionList: function () {
      var list = [{ text: "Feedback-Carousel", color: "red" }];
      if (this.questions.length > 0) {
        for (let i = 0; i < this.members; i++) {
          var q = {};
          var object =
            this.questions[Math.floor(Math.random() * this.questions.length)];
          console.log(object);
          q.text = object.text;
          q.color = this.colors[Math.floor(Math.random() * this.colors.length)];
          list[i + 1] = q;
        }
      }
      return list;
    },
  },
  data() {
    return {
      questions: [],
      model: 0,
      members: 6,
      colors: ["primary", "secondary", "yellow darken-2", "red", "orange"],
    };
  },
  mounted() {
    axios
      .get("http://localhost/api/v1/question")
      .then((res) => {
        this.questions = res.data;
        console.log(this.questions.length + " Question loaded");
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>