//Review all sorts algoritm written in Javascript

var s = "we are happy"

console.log(s);

replaceSpace(s);

// https://www.nowcoder.com/practice/4060ac7e3e404ad1a894ef3e17650423?tpId=13&tqId=11155&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
function replaceSpace(str){
	var final = "";
	for(let s in str){
		if(s === " "){
			s = "%20";

		}
					console.log(s);

		final = final + s
	}
	console.log(final);
}
