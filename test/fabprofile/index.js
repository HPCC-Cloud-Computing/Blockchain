var express = require("express");
var app = express();



var bodyParser = require('body-parser');
var urlencodedParser = bodyParser.urlencoded({ extended: false });
var thongbao;

const program = require('commander');

'use strict';


// cau hinh ejs
app.set("view engine", "ejs");
app.set("views", "./views");

var defaultConfig = require("./config");
var path = require('path');



var store_path = path.join(__dirname, 'hfc-key-store');



/////////////////////////////////////////////////
app.use(express.static('public'));


const config = Object.assign({}, defaultConfig, {
    channelName: "mychannel",
    user: "user1",
    storePath: store_path

});

console.log("Config:", config);

var controller = require("./controller")(config);

var request = {
    //targets : --- letting this default to the peers assigned to the channel
    chaincodeId: "",
    fcn: "",
    args: ['']
};

// each method require different certificate of user
app.get("/home", function (req, res) {
    
    var hung1;
    var id = req.query.userid;
    var lop = req.query.class;
    console.log("lop: ", lop);
    console.log("id: ", id);

    if (typeof id !== "undefined") {

        request.chaincodeId = "aaa1";
        request.fcn = "getUserByID";
        request.args[0] = id;
        console.log(request);
        controller
            .query("user1", request)
            .then(ret => {
                // console.log( "Query results 23131: ",JSON.parse(ret.toString())[0]);

                checkobj = JSON.parse(ret.toString())[0];
                if (typeof checkobj !== "undefined") {
                    hung1 = checkobj.Record;
                    console.log("hung1: ", hung1);
                    
                    if (typeof lop !== "undefined") {
                        request.chaincodeId = "aaa2";
                        request.fcn = "getProfileByID";
                        request.args[0] = id;

                        console.log(request);

                        controller
                            .query("user1", request)
                            .then(ret => {
                                // console.log( "Query results 23131: ",JSON.parse(ret.toString())[0]);

                                userclass = JSON.parse(ret.toString())[0].Record;
                                console.log("userclass= ", userclass);
                                if (lop == "10") {
                                    classid = userclass.class_10;
                                    console.log("Lop 10: ", classid);
                                } else if (lop == "11") {
                                    classid = userclass.class_11;
                                    console.log("Lop 11: ", classid);
                                } else if (lop == "12") {
                                    classid = userclass.class_12;
                                    console.log("Lop 12: ", classid);
                                }
                                console.log("hung: ", hung1);
                               
                                res.render("student/home", { classid,hung: hung1,userclass });

                            })
                            .catch(err => {
                                console.error(err);
                                thongbao="Khong co thong tin ve hoc ba";
                                res.render("student/home", {hung: hung1, thongbao });

                            });
                    } else {
                        res.render("student/home", { hung: hung1 });
                    }
                } else {
                    console.log("Loi khong tim thay");
                    res.render("404_notfound")
                }
            })
            .catch(err => {
                console.error(err);
            });


    }
    else {
        res.render("student/home", { hung: [] });
    }
});

app.get("/lecturer", function(req, res){
    res.render("lecturer/index");
});
app.get("/school", function(req, res){
    res.render("school/index");
});
app.get("/student", function(req, res){
    res.render("student/index");
});

app.post("/notify/:id",urlencodedParser ,function(req, res){

    var id = req.params.id;
    console.log("id: ",id);
    var user_profile= [];
    var stringInput=[];
    var school_pf_tag=["userid","lop","truong","namhoc","hieutruong","gvcn","toan","ly","hanhkiem","danhhieu","bangcap"];

    for(var i=0; i<school_pf_tag.length; i++){
        
        var sp = school_pf_tag[i];
        user_profile[i]=req.body[sp];
        console.log("ok test: ", user_profile);
    }
    if(id == 10){
        stringInput=["aaa2",user_profile[0],user_profile[1]+","+user_profile[2]+","+user_profile[3]+","+user_profile[4]+","+user_profile[5]+","+"Toan#"+user_profile[6]+"&Ly#"+user_profile[7]+","+user_profile[8]+","+user_profile[9],user_profile[10]];
        console.log("string input", stringInput);
        // each method require different certificate of user


        request.chaincodeId = "aaa";
        request.fcn = "initProfile";
        request.args = stringInput;

        controller
        .invoke("user1", request)
        .then(results => {
            console.log(
                "Send transaction promise and event listener promise have completed",
                results
            );
            res.render("notify", {thongbao : "Success"});
            
        })
        .catch(err => {
            console.error(err);
            res.render("notify", {thongbao : "Failed"});
        });

    } else if (i==11 || i==12){
        stringInput=["aaa2",user_profile[0],user_profile[1]+","+user_profile[2]+","+user_profile[3]+","+user_profile[4]+","+user_profile[5]+","+"Toan#"+user_profile[6]+"&Ly#"+user_profile[7]+","+user_profile[8]+","+user_profile[9],id];
        console.log("string input", stringInput);
        // each method require different certificate of user


        request.chaincodeId = "aaa";
        request.fcn = "updateProfile";
        request.args = stringInput;

        controller
        .invoke("user1", request)
        .then(results => {
            console.log(
                "Send transaction promise and event listener promise have completed",
                results
            );
            res.render("notify", {thongbao : "Success"});
        })
        .catch(err => {
            console.error(err);
            res.render("notify", {thongbao : "Failed"});
        });

    }
    
});

app.get("/create/class10",urlencodedParser ,function(req, res){
    res.render("lecturer/createprofile10");
});
app.get("/create/class11",urlencodedParser ,function(req, res){
    res.render("lecturer/createprofile11");
});
app.get("/create/class12",urlencodedParser ,function(req, res){
    res.render("lecturer/createprofile12");
});
app.get("/createstudent",function(req, res){
    res.render("school/student");
});
app.post("/notifystudent",urlencodedParser ,function(req, res){
    var user_inf=["user_id", "name_user", "date_of_brith", "sex_user", "address_user"]
    var user= ["aaa1"];

    for(var i=0; i<user_inf.length; i++){
        
        var sp = user_inf[i];
        user[i+1]=req.body[sp];
        console.log("ok test: ", user);
    }
    console.log("string input", user);
    // each method require different certificate of user


    request.chaincodeId = "aaa";
    request.fcn = "initUser";
    request.args = user;

    controller
        .invoke("user1", request)
        .then(results => {
            console.log(
                "Send transaction promise and event listener promise have completed",
                results
            );
            res.render("notify", {thongbao : "Success"});
        })
        .catch(err => {
            console.error(err);
            res.render("notify", {thongbao : "Failed"});
        });
    

});

app.get("/updateuser", function(req, res){
    var student;
    var id = req.query.userid;
    console.log("id: ", id);

    if (typeof id !== "undefined") {

        request.chaincodeId = "aaa";
        request.fcn = "getUserByID";
        request.args = ["aaa1",id];
        console.log(request);
        
        controller
        .query("user1", request)
        .then(ret => {
            console.log( "Query results 23131: ",JSON.parse(ret.toString())[0]);

            checkobj = JSON.parse(ret.toString())[0];
            if (typeof checkobj !== "undefined") {
                student = checkobj.Record;
                console.log("student: ", student);

                res.render("school/update_user",{ student : student});
            } else {
                console.log("Loi khong tim thay");
                res.render("404_notfound")
            }
        })
        .catch(err => {
            console.error(err);
        });
      


    }
    else {
        res.render("school/update_user");
    }
});

app.post("/notifyuser", urlencodedParser, function(req, res){
    var user_inf=["user_id", "name_user", "date_of_brith", "sex_user", "address_user"]
    var user= ["aaa1"];

    for(var i=0; i<user_inf.length; i++){
        
        var sp = user_inf[i];
        user[i+1]=req.body[sp];
        console.log("ok test: ", user);
    }
    console.log("string input", user);
    // each method require different certificate of user


    request.chaincodeId = "aaa";
    request.fcn = "updateUser";
    request.args = user;

    controller
        .invoke("user1", request)
        .then(results => {
            console.log(
                "Send transaction promise and event listener promise have completed",
                results
            );
            res.render("notify", {thongbao : "Success"});
        })
        .catch(err => {
            console.error(err);
            res.render("notify", {thongbao : "Failed"});

        });
    
    
});
app.get("/deletestudent", function(req, res){
    var id = req.query.userid;
    console.log("id: ", id);
    if(typeof id !== "undefined"){
        request.chaincodeId = "aaa";
        request.fcn = "deleteUser";
        request.args = ["aaa1","aaa2",id];
     console.log("request: ", request);
        controller
            .invoke("user1", request)
            .then(results => {
                console.log(
                    "Send transaction promise and event listener promise have completed",
                    results
                );
                res.render("notify", {thongbao : "Success"});

            })
            .catch(err => {
                console.error(err);
                res.render("notify", {thongbao : "Failed"});

            });
        
    } else {
        res.render("school/delete_student");
    }

});

app.get("/sogd", function(req, res){
    res.render("sogd/index");
})
app.post("/createscoreexam", urlencodedParser, function(req, res){
    var monid=["user_id","mon1","mon2","mon3","mon4"];
    var input_score=[];
    var arg_score=[];
    for(var i=0; i<monid.length; i++){
        var sp = monid[i];
        input_score[i]=req.body[sp];
        console.log("ok input score: ", input_score);
    }
    
    arg_score=["aaa3",input_score[0],monid[1]+"#"+input_score[1]+"&"+monid[2]+"#"+input_score[2]+"&"+monid[3]+"#"+input_score[3]+"&"+monid[4]+"#"+input_score[4]];
    // each method require different certificate of user


    request.chaincodeId = "aaa";
    request.fcn = "initScore";
    request.args = arg_score;
    console.log("request: ",request);
    controller
        .invoke("user1", request)
        .then(results => {
            console.log(
                "Send transaction promise and event listener promise have completed",
                results
            );
            res.render("notify", {thongbao : "Success"});

        })
        .catch(err => {
            console.error(err);
            res.render("notify", {thongbao : "Failed"});

        });  
});

app.get("/createscore", function(req, res){
    res.render("sogd/create_score");
});

app.get("/scoreexam", function(req, res){
    var student;
    var id = req.query.userid;
    console.log("id: ", id);

    if ((typeof id) !== "undefined") {

        request.chaincodeId = "aaa";
        request.fcn = "getScoreByID";
        request.args = ["aaa3",id];
        console.log(request);
        
        controller
        .query("user1", request)
        .then(ret => {
            console.log( "Query results 23131: ",JSON.parse(ret.toString())[0]);

            checkobj = JSON.parse(ret.toString())[0];
            console.log("checkobj: ", checkobj);
            if (typeof checkobj !== "undefined") {
                student = checkobj.Record;
                console.log("student: ", student);

                res.render("student/get_score", {student : student});
            } else {
                console.log("Loi khong tim thay");
                res.render("404_notfound")
            }
        })
        .catch(err => {
            console.error(err);
        });

    }
    else {
        res.render("student/get_score");
    }
});

app.get("/notifygrad",function(req, res){
    var student;
    var id = req.query.userid;
    console.log("id: ", id);

    if ((typeof id) !== "undefined") {

        request.chaincodeId = "aaa";
        request.fcn = "checkScore";
        request.args = ["aaa1","aaa2","aaa3",id];
        console.log(request);
        
        controller
        .query("user1", request)
        .then(ret => {
            console.log( "Query results 23131: ",JSON.parse(ret.toString()));

            checkobj = JSON.parse(ret.toString());
            console.log("checkobj: ", checkobj);
            if (typeof checkobj !== "undefined") {
               

                res.render("student/get_score", {bangcap : checkobj});
            } else {
                console.log("Loi khong tim thay");
                res.render("404_notfound")
            }
        })
        .catch(err => {
            console.error(err);
        });

    }
    else {
        res.render("student/get_score");
    }
});

app.listen(4200);