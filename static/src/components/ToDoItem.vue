<!-- 
    单文件组件：
    位于components目录下
    在app.vue中注册
 -->
 <!-- 
    prop -- 组件间通信
    1.单向数据传递：
        父组件向子组件传递数据
  -->
<template>
    <div v-if="!isEditing">
            <!-- 
                现在 我们可以做到动态渲染子组件了
                接下来，我们将实现 页面读取用户输入，并加载输入数据到页面中
                因此我们需要：
                    1. 一个<input> 标签
                    2. 一个方法，用于实现获取输入并展示出来
                    3. 一个控制数据的模型

            -->
        <div>
           <input 
            type="checkbox" 
            name="my-first-checkbox" 
            :id="id" 
            :checked="this.isDone" 
            @change="$emit('checkbox-changed')"/>
            <label>{{ label }}</label> 
        </div>
        <div class="">
            <button 
                type="button" 
                @click="toggleToItemEditForm"
                ref="editButton">
                Edit 
                <!-- <span class="visually-hidden">{{label}}</span> -->
            </button>
            <button type="button" @click="deleteToDo">
                Delete 
                <!-- <span class="visually-hidden">{{label}}</span> -->
            </button>
        </div>
    </div>
    <to-do-item-edit-form 
        v-else :id="id" 
        :label="label"
        @item-edited="itemEdited"
        @edit-cancelled="itemCancelled">
        </to-do-item-edit-form>
</template>

<script>
import ToDoItemEditForm from './ToDoItemEditForm.vue';
export default {
    components: {
        ToDoItemEditForm
    },
    props:{
        label: {
            required: true,
            type: String,
        },
        done: {
            required: true,
            type: Boolean
        },
        id: {
            required: true,
            type: String,
        }
    },
    data()  {
        return {
            
            isEditing: false,
        };
    },
    methods: {
        deleteToDo(){
            this.$emit('item-deleted');
        },
        toggleToItemEditForm() {
            console.log(this.$refs.editButton);
            this.isEditing = true;
            
        },
        itemEdited(newlabel) {
            this.$emit('item-edited', newlabel);
            this.isEditing=false;
            this.focusOnEditButton();
        },
        itemCancelled() {
            this.isEditing = false;
            this.focusOnEditButton();
        },
        focusOnEditButton() {
            this.$nextTick(() => {
                const editButtonRef = this.$refs.editButton;
                editButtonRef.focus();
            });
        }
    },
    computed: {
        isDone() {
            return this.done
        }
    }
};
</script>


