<script setup>
import ToDoItem from './components/ToDoItem.vue';
import uniqueId from "lodash.uniqueid";
import ToDoForm from './components/ToDoForm.vue';
import axios from 'axios';
</script>

<script>
export default {
  created() {
    //请求后端 拉取数据
    
  },
  mounted() {
    this.fetchData()
  },
  data() {
    return {
      ToDoItems: [

      ],
    };
  },
  methods: {
    addToDo(toDoLabel) {
      //console.log("To-Do added", toDoLabel);
      this.ToDoItems.push({
        id: uniqueId(),
        label: toDoLabel,
        done: false
      });
    },
    updateDoneStatus(toDoId) {
      //概念 ： js箭头函数
      //概念 ： js 数组的 find filter方法 
      const toDoUpdate = this.ToDoItems.find((i) => i.id ===toDoId)
      toDoUpdate.done = !toDoUpdate.done
    },
    editToDo(newlabel,id) {
      this.ToDoItems.find(i => i.id === id).label = newlabel;
    },
    deleteToDo(toDoId) {
      const itemIndex = this.ToDoItems.findIndex((item) => item.id === toDoId);
      this.ToDoItems.splice(itemIndex, 1);
      this.$refs.listSummary.focus();
    },
    fetchData(){
      axios.get('/getAllToDoItems')
      .then( res => {
        //console.log(res);
        var data = res.data.data;
        this.ToDoItems = data;
      })
      .catch(err => {

      })
      .then( end => {

      });
    },
  },
  computed: {
    //cmoputed属性中的变量依赖于其他变量，vue能够感知totalQuantity所依赖的变量的变化，从而实现懒加载
    listSummary() {
      const numberFinishedItems = this.ToDoItems.filter((item) => item.done).length
      // `${}` 是ES6引入的新写法，方便了许多
      return `${this.ToDoItems.filter((item) => item.done).length} out of ${this.ToDoItems.length} items completed`
    }
  },
};

</script>
<template>
  <header>
    <div id="app">
      <h1>To-Do List</h1>

      <!-- 
        事件监听器 todo-added 事件触发时将调用 addToDo()方法 
        @后接 要监听的事件 ""内表示要调用的函数名
      -->
      <to-do-form @todo-added="addToDo">

      </to-do-form>
      <h2 ref="listSummary" tabindex="-1" id="list-summary">{{listSummary}}</h2>
      <ul>
        <!-- 
          v-for 是Vue.js中的一个指令，用于循环遍历数组或对象的元素，并生成相应的DOM元素。
          :key 不是必须的，但是它有利于vue更加高效
         -->
        <li v-for="item in ToDoItems" :key="item.id">
          <!-- 
            :id="item.id" 是将 item.id 传递到子组件 <to-do-item> 中的 prop 属性中，
            而不是直接设置HTML元素的 id 属性。
            同理 @checkbox-changed 也是在监听子组件中 checkbox-changed事件的发生
           -->
          <to-do-item 
          
            @checkbox-changed="updateDoneStatus(item.id)"
            @item-edited="editToDo($event,item.id)"
            @item-deleted="deleteToDo(item.id)"
            :label="item.label" 
            :id="item.id"
            :done="item.done"></to-do-item>
        </li>
      </ul>
    </div>
  </header>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
