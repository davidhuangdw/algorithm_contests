package google_kickstart.`2020`.H

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N,K,S) = input_ints()
        println("Case #${it}: ${N + minOf(K, 2*(K-S))}")
    }
}
