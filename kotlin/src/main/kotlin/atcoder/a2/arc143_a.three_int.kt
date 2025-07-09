package atcoder.a2

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (a, b, c) = input_longs().sorted()
    if(c <= a+b){
        println(c)
    }else{
        println(-1)
    }
}