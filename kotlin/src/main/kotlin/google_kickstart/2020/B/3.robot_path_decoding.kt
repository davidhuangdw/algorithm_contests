package google_kickstart.`2020`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val MOD = 1e9.toLong()
    val T = readLine()!!.toInt()
    (1..T).forEach {
        var (x, y) = 0L to 0L
        val pos_st = mutableListOf<Pair<Long, Long>>()
        val times_st = mutableListOf<Long>()

        var times = 0L
        for(ch in readLine()!!){
            if(ch.isDigit()){
                times = (times*10 + (ch - '0').toLong()) % MOD
            }else when(ch){
                '(' -> {
                    pos_st.add(x to y)
                    times_st.add(times)
                    x = 0
                    y = 0
                    times = 0L
                }
                ')' -> {
                    times = times_st.removeLast()
                    x = x*times % MOD
                    y = y*times % MOD
                    val (px, py) = pos_st.removeLast()
                    x = (px + x) % MOD
                    y = (py + y) % MOD
                    times = 0L
                }
                'S' -> y = (y+1) % MOD
                'N' -> y = (y-1+MOD) % MOD
                'E' -> x = (x+1) % MOD
                'W' -> x = (x-1+MOD) % MOD
            }
        }
        x += 1
        y += 1
        println("Case #${it}: $x $y")
    }
}
