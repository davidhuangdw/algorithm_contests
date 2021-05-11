package code.jam.`2020`.qua
// https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/0000000000209a9f#problem

fun paired(S: String): String{
    return buildString {
        var pre = 0
        for(ch in S){
            val k = ch - '0'
            repeat(pre - k){append(")")}
            repeat(k - pre){append('(')}
            append(ch)
            pre = k
        }
        repeat(pre){append(')')}
    }
}

fun main(){
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val S = readLine()!!
        println("Case #${it}: ${paired(S)}")
    }
}