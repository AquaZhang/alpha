<template>
    <!-- @envent 的属性
          @event.modifier
            .stop：停止传播事件。等效于常规 JavaScript 事件中的 Event.stopPropagation()。
            .prevent：阻止事件的默认行为。等效于 Event.preventDefault()。
            .self：仅当事件是从该确切元素分派时触发处理程序。
            {.key}：仅通过指定键触发事件处理程序。 MDN 有一个有效键值列表 (en-US); 多词键只需转换为 kebab 大小写（例如 page-down）。
            .native：监听组件根（最外层的包装）元素上的原生事件。
            .once：监听事件，直到它被触发一次，然后不再触发。
            .left：仅通过鼠标左键事件触发处理程序。
            .right：仅通过鼠标右键事件触发处理程序。
            .middle：仅通过鼠标中键事件触发处理程序。
            .passive：等效于在 vanilla JavaScript 中使用 addEventListener() 创建事件监听器时传入 { passive: true } 参数。
     -->
    <form @submit.prevent="onSubmit">
        <label for="new-todo-input">What needs to be done?</label>
        <input 
            type="text" 
            id="new-todo-input"
            name="new-todo"
            autocomplete="off" 
            v-model.lazy.trim="label"/>
            <!-- 
                v-model:
                在这个例子中,input的元素将与label双向绑定
             -->
        <button type="submit">ADD</button>
    </form>
</template>

<script>
    
    import axios from 'axios';
    export default{
        methods: {
            onSubmit(){
                console.log("form submitted");
                console.log('label value: ', this.label);
                //添加一个自定义事件 todo-added (区分大小写，不能包含空格)
                if(this.label === "") {
                    return;
                }
                this.$emit("todo-added",this.label);
                this.label = "";
                
                //新加网络请求
                axios.get('/hello')
                .then( res => {
                    //成功
                    console.log("success")
                    console.log(res)
                })
                .catch( err => {
                    //错误处理
                })
                .then( end => {

                })

            },
        },
        //添加一个data()方法
        data() {
            return {
                label: "",
            };
        },
    };
</script>