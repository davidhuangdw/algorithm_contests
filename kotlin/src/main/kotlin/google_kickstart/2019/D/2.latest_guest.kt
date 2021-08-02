package google_kickstart.`2019`.D

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, G, M) = input_ints()
        val st = (0 until G).map {
            val (x, d) = readLine()!!.split(" ")
            x.toInt() - 1 to if (d == "C") 1 else -1
        }

        fun count(ids: List<List<Int>>): List<Int> {
            val cnt = MutableList(G) { 0 }
            for (l in ids)
                for (g in l)
                    cnt[g] += 1
            return cnt
        }

        fun testset1(): List<Int> {
            val last = hashMapOf<Int, MutableList<Int>>()
            repeat(M + 1) { m ->
                val cur = hashMapOf<Int, MutableList<Int>>()
                for ((i, g) in st.withIndex()) {
                    val (x, d) = g
                    val y = ((x + 1L * (N + d) * m) % N).toInt()
                    if (y !in cur)
                        cur[y] = mutableListOf()
                    cur[y]!!.add(i)
                }
                for ((x, l) in cur)
                    last[x] = l
            }
            return count(last.values.toList())
        }

        fun testset2(): List<Int> {
            // optimize idea:
            //      group the guests in the same start_points,
            //      because if K guests are same group then count would be O(Consultants * K) = O(N * G)
            //      count on groups, then O(consultants) to get pair(group_id, cnt), then O(G) to apply on each guest
            fun toGroups(guests: HashSet<Int>, d: Int): HashMap<Int, MutableList<Int>> {
                val groups = hashMapOf<Int, MutableList<Int>>()
                for (g in guests) {
                    val x = st[g].first
                    if (x !in groups)
                        groups[x] = mutableListOf()
                    groups[x]!!.add(g)
                }
                return groups
            }

            fun calc(groups: HashSet<Int>, d: Int): HashMap<Int, Pair<Int, Int>> {
                val street = hashMapOf<Int, Pair<Int, Int>>()
                val remain = groups
                for (m in M downTo 0) {
                    val del = hashSetOf<Int>()
                    for (g in remain) {
                        val pos = ((g + 1L * (N + d) * m) % N).toInt()
                        if (pos in street) {
                            del.add(g)
                            continue
                        }
                        street[pos] = m to g
                    }
                    remain -= del
                    if (remain.isEmpty()) break
                }
                return street
            }

            val ga = hashSetOf<Int>()
            val gb = hashSetOf<Int>()
            for ((g, v) in st.withIndex())
                if (v.second > 0) ga.add(g)
                else gb.add(g)
            val gra = toGroups(ga, 1)
            val grb = toGroups(gb, -1)
            val sa = calc(gra.keys.toHashSet(), 1)
            val sb = calc(grb.keys.toHashSet(), -1)

            // union
            val street = hashMapOf<Int, List<Int>>()
            val ac = hashMapOf<Int, Int>()
            val bc = hashMapOf<Int, Int>()
            for ((pos, va) in sa) {
                val vb = sb.getOrDefault(pos, -1 to -1)
                if (va.first >= vb.first)
                    ac[va.second] = ac.getOrDefault(va.second, 0) + 1
                if (va.first <= vb.first)
                    bc[vb.second] = bc.getOrDefault(vb.second, 0) + 1
            }
//            println(street)
            for ((pos, vb) in sb)
                if (pos !in sa)
                    bc[vb.second] = bc.getOrDefault(vb.second, 0) + 1

            val count = MutableList(G) { 0 }
            for ((group, cnt) in listOf(gra to ac, grb to bc))
                for ((gr, c) in cnt)
                    for (g in group[gr]!!)
                        count[g] += c
            return count
        }

        val res = testset2().joinToString(" ")
        println("Case #${it}: $res")
    }
}
