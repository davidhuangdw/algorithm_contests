package atcoder.a8

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val (T, ) = input_ints()
    repeat(T){
        val (N, ) = input_ints()
        repeat(3){readLine()}
        println((listOf(1) + List(N){0} + List(N){1}).joinToString(""))
    }
}