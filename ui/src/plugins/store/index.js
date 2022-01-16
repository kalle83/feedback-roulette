// import vue and vuex
import axios from 'axios'
// import { Promise } from 'core-js'
import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)


// create our store
export default new Vuex.Store({
  state: {
    apiEndpoint: "http://localhost",
    fotterMessage: "Regardless of what we discover, we understand and truly believe that everyone did the best job they could, given what they knew at the time, their skills and abilities, the resources available, and the situation at hand.",
    questions: [],
    isLoading: false,
  },

  mutations: {
    UPDATE_QUESTIONS: (state, questions) => {
      state.questions = questions
    },
    START_LOADING: (state) => {
      state.isLoading = true
    },
    END_LOADING: (state) => {
      state.isLoading = false
    }

  },

  actions: {
    loadQuestions(context) {
      return new Promise((resolve, reject) => {
        context.commit("START_LOADING")
        console.time("LOAD-QUESTIONS")
        axios({
          baseURL: context.state.apiEndpoint,
          url: '/api/v1/question',
          method: 'GET'
        }).then((response) => {
          if (response.status == 200) {
            context.commit('UPDATE_QUESTIONS', response.data)
            console.log(response.data)
            resolve()
          }
        }).catch((error) => {
          if (error.response) {
            console.log("Response Error: " + error.response.status + " " + error.response.statusText);
          } else if (error.request) {
            console.log(error.request);
          } else {
            console.log('Error', error.message);
          }
          reject(error)
        }).finally(() => {
          context.commit('END_LOADING')
          console.timeEnd("LOAD-QUESTIONS")
        })
      })
    }
  },


  getters: {
    getFooterMsg(state) {
      return state.fotterMessage
    }
  },
})