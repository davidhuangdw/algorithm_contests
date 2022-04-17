package google_code_jam.`2008`.qua

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (S) = input_ints()
        val Ss = (1..S).map { split_input()[0] }.toHashSet()
        val (Q) = input_ints()

        val vis = hashSetOf<String>()
        var cnt = 0
        (1..Q).forEach {
            val q = readLine()!!.trim()
            vis.add(q)
            if(vis.size == S){
                cnt ++
                vis.clear()
                vis.add(q)
            }
        }

        println("Case #${it}: $cnt")
    }
}
