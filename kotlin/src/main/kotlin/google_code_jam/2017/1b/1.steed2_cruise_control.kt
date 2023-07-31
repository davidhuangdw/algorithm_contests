package google_code_jam.`2017`.`1b`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (D, N) = input_ints()
        val A = (1..N).map{input_ints()}
        var t = A.maxOfOrNull{(D-it[0])*1.0/it[1]}!!
        var r = D/t
        println("Case #${it}: $r")
    }
}
