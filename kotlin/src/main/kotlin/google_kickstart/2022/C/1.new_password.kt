package google_kickstart.`2022`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (N) = input_ints()
        val S = split_input()[0]
        var up = false
        var low = false
        var dig = false
        var sp = false
        for (ch in S) {
            if (ch in '0'..'9') {
                dig = true
            } else if (ch in 'a'..'z')
                low = true
            else if (ch in 'A'..'Z')
                up = true
            else sp = true
        }
        val add = mutableListOf<Char>()
        if (!up)
            add.add('A')
        if (!low)
            add.add('a')
        if (!dig)
            add.add('1')
        if (!sp)
            add.add('@')
        for (i in 0 until 7 - (N + add.size)) {
            add.add('1')
        }
        val r = S + add.joinToString("")
        println("Case #${it}: $r")
    }
}
