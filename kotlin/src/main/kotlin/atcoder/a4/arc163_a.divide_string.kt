package atcoder.a4

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    val (T, ) = input_ints()
    repeat(T){
        readLine()!!
        val S = readLine()!!
        val n = S.length
        println(if((1 until n).any{S.substring(0, it) < S.substring(it)}) "Yes" else "No")
    }
}