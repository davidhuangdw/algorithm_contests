package atcoder.a8

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_longs() = split_input().map { it.toLong() }

    val (N, L, R) = input_longs()
    var sum = 0L
    var rem = N
    while(rem > 0L){
        val lowbit = rem and -rem
        val l = lowbit
        val r = (lowbit shl 1) - 1
        sum += maxOf(0L, minOf(r, R) - maxOf(l, L)+1)
//        println(listOf(rem, lowbit, l, r))
        rem -= lowbit
    }
    println(sum)
}