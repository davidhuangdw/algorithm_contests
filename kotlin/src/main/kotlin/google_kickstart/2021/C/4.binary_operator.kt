package google_kickstart.`2021`.C

import java.math.BigInteger

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        var units = hashMapOf(listOf(1) to 0)
        var count = 1
        val (N, ) = input_ints()
        repeat(N){
            val S = readLine()!!
            var i = 0
            val ops = mutableListOf<Char>()
            val result = mutableListOf<Pair<BigInteger, Char>>()
            val priority = hashMapOf('(' to -1, '+' to 0, '*' to 1, '#' to 2)
            while(i < S.length){
                if(S[i].isDigit()){
                    val j = i
                    while(i < S.length && S[i].isDigit()) i+=1
                    result.add(S.substring(j, i).toBigInteger() to 'x')
                }else if(S[i] == '('){
                    ops.add(S[i])
                    i+= 1
                }else if(S[i] == ')'){
                    while(ops.isNotEmpty() && ops.last() != '('){
                        result.add(0.toBigInteger() to ops.removeLast())
                    }

                    i+=1
                }else{
                    while(ops.isNotEmpty() && priority[ops.last()]!! >= priority[S[i]]!!){
                        result.add(0.toBigInteger() to ops.removeLast())
                    }
                    ops.add(S[i])
                    i+=1
                }

            }
        }
        println("Case #${it}: ")
    }
}
