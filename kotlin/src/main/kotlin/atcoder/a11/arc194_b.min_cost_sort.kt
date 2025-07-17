package atcoder.a11

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
    val (N) = inputInts()
    val P = inputInts()

    var a = P.toIntArray()
    val cnt = LongArray(N+1)

    fun merge(i: Int, len: Int){
        if(len <= 1) return
        val l = len/2
        val md = i+l
        merge(i, l)
        merge(md, len-l)
        var cur = mutableListOf<Int>()
        var k = md
        for(j in i until md){
            var x = a[j]
            while(k < i+len && a[k] < x){
                cur.add(a[k])
                k ++
            }
            cnt[x] += 0L+k-md
            cur.add(x)
        }
        for(j in k until i+len){
            cur.add(a[j])
        }
        for(j in 0 until len){
            a[i+j] = cur[j]
        }
    }
    merge(0, N)
    println((1..N).sumOf{
        val k = cnt[it]
        (2*it - k - 1)*k/2
    })
}
