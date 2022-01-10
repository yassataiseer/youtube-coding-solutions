const app = require('express')();

const PORT = 8080

app.get("/grab_Users",(req,res)=>{
    res.status(200).send({
        name:"Yassa",
        id:id
    })

});


app.listen(PORT,()=>console.log(`VIEW API ON http://localhost:${PORT}`))