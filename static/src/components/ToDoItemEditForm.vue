<template>
    <form class="" @submit.prevent="onSubmit">
        <div>
            <label for="">Edit :  
            </label>
            <input 
                type="text" 
                name=""
                ref="labelInput"
                id=""
                autocomplete="off"
                v-model.lazy.trim="newlabel" />
        </div>
        <div class="">
            <button type="button" class="" @click="onCancel">
                Cancel
                <!-- <span class="visually-hidden">editing {{ label }}</span> -->
            </button>
            <button type="button" class="" @click="onSubmit">
                Save
                <!-- <span class="visually-hidden">edit for {{ label }}</span> -->
            </button>
        </div>
    </form>
</template>

<script>
export default {
    props: {
        label: {
            type: String,
            required: true,
        },
        id: {
            type: String,
            required: true,
        },
    },
    data() {
        return {
            newlabel: this.label,
        };
    },
    methods: {
        onSubmit() {
            if(this.newlabel && this.newlabel !== this.label) {  
                this.$emit("item-edited", this.newlabel);
            }
        },
        onCancel() {
            this.$emit("edit-cancelled");
        },
    },
    mounted() {
        const labelInputRef = this.$refs.labelInput;
        labelInputRef.focus();
    },
}
</script>

<style>

</style>