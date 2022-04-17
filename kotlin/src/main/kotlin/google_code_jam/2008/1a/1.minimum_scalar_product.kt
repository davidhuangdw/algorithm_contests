package google_code_jam.`2008`.`1a`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val X = input_longs().sorted()
        val Y = input_longs().sortedDescending()
        val sum =X.zip(Y).sumOf{(x, y) -> x*y}
        println("Case #${it}: $sum")
    }
}
