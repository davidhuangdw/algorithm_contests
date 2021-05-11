package code.jam.`2020`.qua


fun main(){
    val T = readLine()!!.toInt()
    repeat(T) {
        val N = readLine()!!.toInt()
        val rows = (0 until N).map{ hashMapOf<Int, Int>() }
        val cols = (0 until N).map{ hashMapOf<Int, Int>() }
        var trace = 0
        for(i in 0 until N){
            val row = readLine()!!.trim().split(' ').map{ it.toInt() }
            for(j in 0 until N){
                val v = row[j]
                rows[i].set(v, rows[i].getOrDefault(v, 0)+1)
                cols[j].set(v, cols[j].getOrDefault(v, 0)+1)
                if(i == j) trace += v
            }
        }
        val repeat_rows = rows.count{ it.any{(k, v) -> v > 1  } }
        val repeat_cols = cols.count{ it.any{(k, v) -> v > 1  } }

        println("Case #${it+1}: $trace $repeat_rows $repeat_cols")
    }
}
