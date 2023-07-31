package google_code_jam.`2019`.`1b`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (P, Q) = input_ints()
        val dir = hashMapOf<Char, MutableList<Int>>()
        for(d in "NSWE")
            dir[d] = MutableList(Q+1){0}
        for(_i in 0 until P){
            val (_x, _y, _d) = split_input()
            val (x, y) = _x.toInt() to _y.toInt()
            val d = _d[0]
            val v = when(d){
                'N' -> y+1
                'S' -> y-1
                'E' -> x+1
                'W' -> x-1
                else -> -1
            }
            dir[d]!![v] = dir[d]!![v] + 1
        }
        fun find(ge: MutableList<Int>, le: MutableList<Int>): Int{
            var cur = le.sum() + ge[0]
            var mx = 0
            var max = cur
            for(x in 1..Q){
                cur += ge[x] - le[x-1]
                if(cur > max){
                    max = cur
                    mx = x
                }
            }
            return mx
        }
        val x = find(dir['E']!!, dir['W']!!)
        val y = find(dir['N']!!, dir['S']!!)
        println("Case #${it}: $x $y")
    }
}
