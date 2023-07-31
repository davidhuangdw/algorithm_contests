package google_code_jam.`2018`.`1c`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, ) = input_ints()
        val cnt = MutableList(N) { 0 }
        val done = MutableList(N) { false }
        for (i in 0 until N) {
            var flavor = input_ints()
            if (flavor[0] == -1)
                throw RuntimeException("failed: $it: $i: $flavor")
            flavor = flavor.subList(1, flavor.size)
            for (x in flavor) {
                cnt[x]++
            }
            val x = flavor.filter { !done[it] }.minByOrNull { cnt[it] }
            if(x == null){
                println(-1)
            }else {
                done[x] = true
                println(x)
            }
        }
//        println("Case #${it}: ")
    }
}
