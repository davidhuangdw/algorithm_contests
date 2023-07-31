package google_kickstart.`2022`.G
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (M, N, P) = input_ints()
        val S = (1..M).map { input_ints() }
        val ma = (0 until N).map{i -> S.maxOf { it[i] }}
        val res = (0 until N).sumOf { maxOf(0, ma[it] - S[P-1][it]) }
        println("Case #${it}: $res")
    }
}
