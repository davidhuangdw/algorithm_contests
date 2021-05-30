package google_kickstart.`2020`.H

import java.lang.Math.abs


fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    // min{ sum{y - yi} } = median{y}
    // viewpoint: x0: min{ sum{x0 + i - xi} = sum{x0 - (xi-i)} } = median{xi - i}
    (1..T).forEach {
        val (N,) = input_ints()
        val (X, Y) = LongArray(N) to LongArray(N)
        repeat(N){ i ->
            val (x, y) = input_longs()
            Y[i] = y
            X[i] = x
        }

        X.sort()
        Y.sort()
        for(i in 0 until N) X[i] -= i.toLong()
        var res = 0L
        val md = (N-1)/2
        for(v in Y)
            res += abs(v - Y[md])
        for(v in X)
            res += abs(v - X[md])
        println("Case #${it}: $res")
    }
}
