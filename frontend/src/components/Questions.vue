<template>
  <div>
    <h1>Students and Questions</h1>
    <button class="btn" v-on:click="studentsf">load students</button>
    <button class="btn" v-on:click="questionsf">load questions</button>
    <div class="flex-container" v-if="students">
      <div v-bind:key="student.id" v-for="student of students">
        <h3>{{ student.id }} - {{ student.firstName}} {{student.middleName }} {{ student.lastName }}</h3>
        <p>Role: {{  (student.role === "S") ? ("Student") : ((student.role === "TA") ? ("Hjælpelære") : ("Underviser")) }}</p>
        <p>DiscordID: {{ student.discord }}</p>
      </div>
      <h2>Students</h2>
    </div>
    <div class="flex-container" v-if="questions">
      <div v-bind:key="question.timestamp" v-for="question of questions">
        <h3>{{ question.student.id }} - {{ question.student.firstName }} {{ question.student.middleName }} {{ question.student.lastName }}</h3>
        <p><small>
          {{ question.topic }},
          {{ new Date(question.timestamp*1000).toUTCString() }},
          {{ question.channelId }},
          {{ question.student.discord }}
        </small></p>
        <p>{{ question.question }}</p>
      </div>
      <h2>Questions</h2>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Questions',
  data() {
    return {
      questions: [],
      students: [],
      errors: [],
    }
  },
  methods : {
    questionsf() {
      axios.post(`http://127.0.0.1:8080/`, {
        query: `
        {
          questions {
            active
            timestamp
            channelId
            topic
            question
            student {
              firstName
              middleName
              lastName
              role
              id
              discord
            }
          }
        }
      `
      }).then(res => {
        this.questions = res.data.data.questions
      }).catch(e => {
        console.log(e)
      })
    },
    studentsf() {
      axios.post(`http://127.0.0.1:8080/`, {
        query: `
        {
          students{
            id
            firstName
            middleName
            lastName
            role
            discord
          }
        }
      `
      }).then(res => {
        this.students = res.data.data.students
      }).catch(e => {
        console.log(e)
      })
    }
  }
}
</script>

<style scoped>
.flex-container {
  display: flex;
  flex-direction: column-reverse;
  background-color: dimgrey;
  padding-left: 5px;
  padding-right: 5px;
  padding-top: 5px;
}
.flex-container > div {
  background-color: darkgray;
  margin-bottom: 5px;
  padding: 5px;
}

</style>