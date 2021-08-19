package google_kickstart.`2018`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    val bits_count = MutableList(1 shl 16) { it }
    for (x in 2 until (1 shl 16))
        bits_count[x] = bits_count[x / 2] + (x and 1)

    (1..T).forEach {
        val (NN, KK, P) = input_longs()
        val (N, K) = NN.toInt() to KK.toInt()
        fun testset1(): String {
            val C = hashMapOf<Int, Int>()
            repeat(K) {
                val (a, b, c) = input_ints()
                C[a - 1] = c
            }
            val st = mutableListOf<Long>()
            var x = P - 1
            repeat(N - C.size) {
                st.add(x % 2)
                x /= 2
            }

            return buildString {
                for (i in 0 until N) {
                    if (i in C) append(C[i])
                    else append(st.removeLast())
                }
            }
        }

        fun tesetset2(): String {
            val C = hashMapOf<Int, MutableList<Pair<Int, Int>>>()
            repeat(K) {
                val (a, b, c) = input_ints()
                if (b !in C)
                    C[b] = mutableListOf()
                C[b]!!.add(b - a + 1 to c)
            }
            val M = 1 shl 15

            val vcache = hashMapOf<Pair<Int, Int>, Boolean>()
            fun valid(i: Int, y: Int): Boolean {
                if (i to y !in vcache) {
                    var vd = true
                    for ((k, c) in C.getOrDefault(i, mutableListOf())) {
                        if (c != bits_count[y and ((1 shl k) - 1)]) {
                            vd = false
                            break
                        }
                    }
                    vcache[i to y] = vd
                }
                return vcache[i to y]!!
            }

            val cache = hashMapOf<Pair<Int, Int>, Long>()
            fun dp(i: Int, x: Int): Long {
                if (i > N) return 1L
                if (i to x !in cache) {
                    var cnt = 0L
                    for (b in 0..1) {
                        val y = x shl 1 or b
                        if (valid(i, y))
                            cnt += dp(i + 1, y and (M - 1))
                    }

                    cache[i to x] = cnt
                }
                return cache[i to x]!!
            }

            var p = P
            var x = 0
            return buildString {
                for (i in 1..N) {
                    x = x shl 1 and (M - 1)
                    var done = false
                    for (b in 0..1) {
                        x = x or b
                        if (!valid(i, x))
                            continue
                        if (p <= dp(i + 1, x)) {
                            append(b)
                            done = true
                            break
                        }
                        p -= dp(i + 1, x)
                    }
                    if (!done)
                        throw RuntimeException("impossible!")   // todo: fix RE
                }
            }
        }

        println("Case #${it}: ${tesetset2()}")
    }
}
