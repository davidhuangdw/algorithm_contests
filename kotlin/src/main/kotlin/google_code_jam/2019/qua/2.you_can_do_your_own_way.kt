package google_code_jam.`2019`.qua
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val (A, ) = split_input()
        val res = A.map{if(it == 'E') 'S' else 'E'}.joinToString("")
        println("Case #${it}: $res")
    }
}
