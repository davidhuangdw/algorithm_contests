package google_kickstart.`2022`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun isIntr(x: Int): Boolean {
        var t = x
        var (sum, mul) = 0 to 1
        while (t > 0) {
            val d = t % 10
            sum += d
            mul *= d
            t /= 10
        }
        return mul % sum == 0
    }

    val cnt = MutableList(100_001) { 0 }
    var cnt_init = false
    fun init_cnt() {
        for (x in 1..100_000) {
            cnt[x] = cnt[x - 1] + if (isIntr(x)) 1 else 0
        }
        cnt_init = true
    }

    fun testset1(A: Long, B: Long): Int {
        if (!cnt_init)
            init_cnt()
        return cnt[B.toInt()] - cnt[A.toInt() - 1]
    }

    val POW = MutableList(14) { 1L }
    for (i in 1 until POW.size)
        POW[i] = POW[i - 1] * 10
    data class Ctx(val P: Long, val S: Long, val L: Int)

    val cache = hashMapOf<Ctx, Long>()
    fun countCtx(ctx: Ctx): Long {
        if (ctx !in cache) {
            var (p, s, l) = ctx
            if (s == 0L) // bug: not handle s==0L: leading zero shouldn't affect p
                p = 1
            var res = 0L
            if (l == 0) {
                res = if (s != 0L && p % s == 0L) 1 else 0
            } else {
                for (d in 0..9) {
                    res += countCtx(Ctx(p * d, s + d, l - 1))
                }
            }
            cache[ctx] = res
        }
        return cache[ctx]!!
    }

    fun countWithN(N: Long, ctx: Ctx, eq_pre: Boolean): Long {
        if (N == 100L) {
            println("hi")
        }
        if (eq_pre) {
            val (p, s, l) = ctx
            if (l == 0)
                return countCtx(ctx)
            val nd = ((N / POW[l - 1]) % 10).toInt()
            var res = 0L
            for (d in 0..minOf(nd, 9)) {
                res += countWithN(N, Ctx(p * d, s + d, l - 1), d == nd)
            }
            return res
        } else {
            return countCtx(ctx)
        }
    }

    fun getL(N: Long): Int {
        if (N == 0L) return 1
        for (i in 0..13) {
            if (POW[i] > N)
                return i
        }
        return -1
    }

    fun countN(N: Long): Long {
        val l = getL(N)
        return countWithN(N, Ctx(1, 0, getL(N)), true)
    }

    fun testset2(A: Long, B: Long): Long {
        return countN(B) - countN(A - 1)
    }

//    init_cnt()
//    for(x in 0..1000){
//        val (a, b) = cnt[x] to countN(x.toLong())
//        if(a != b.toInt()){
//            throw RuntimeException("not the same $x: $a $b")
//        }
//    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (A, B) = input_longs()
        println("Case #${it}: ${testset2(A, B)}")
    }
}
