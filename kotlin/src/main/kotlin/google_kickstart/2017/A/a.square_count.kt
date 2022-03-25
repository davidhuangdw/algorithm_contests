package google_kickstart.`2017`.A

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val M = 1000_000_007L
    fun add(vararg vs: Long): Long {
        var x = 0L
        vs.forEach {
            x = (x + it) % M
        }
        return x
    }
    fun mul(vararg vs: Long): Long {
        var x = 1L
        vs.forEach {
            x = x * it % M
        }
        return x
    }
    fun pow(a: Long, p: Long): Long{
        var b = a
        var x = 1L
        var t = p
        while(t > 0) {
            if(t and 1L > 0){
                x = x * b % M
            }
            b = b * b % M
            t /= 2
        }
        return x
    }
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C) = input_longs()
        val N = minOf(R, C)

        val sum = add(
            mul(R, C, N, N+1, pow(2, M-2)), // bug: N*(N+1)/2, might > M, so should be (N*(N+1)/2 % M)
            M - mul(R+C, N, N+1, 2*N+1, pow(6, M-2)),
            mul(N, N, N+1, N+1, pow(4, M-2)))

        println("Case #${it}: $sum")
    }
}
