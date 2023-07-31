package google_code_jam.`2019`.r2

import kotlin.math.absoluteValue

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    fun gcd(_x: Long, _y: Long): Long{
        var (x, y) = _x to _y
        while(x != 0L){
            x = y%x.also{y = x}
        }
        return y
    }
    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, ) = input_ints()
        val P = (1..N).map{input_longs()}

        // idea: order of each pair => unique order of all items (if no contradict)
        // all uncertain pairs put together, their order could be uncertain(at least 1 valid order), depend only on wx/wy
        val ratios = hashSetOf<Pair<Long, Long>>()
        for((i, ip) in P.withIndex())
            for(j in i+1 until N){
                val (xi, yi) = ip
                val (xj, yj) = P[j]
                var (dx, dy) = xi- xj to yi - yj
                if(dx*dy >= 0) continue
                dx = dx.absoluteValue
                dy = dy.absoluteValue
                val g = gcd(dx, dy)
                dx /= g
                dy /= g
                ratios.add(dx to dy)
            }
        val res = ratios.size + 1
        println("Case #${it}: $res")
    }
}
