package main

/*
// ConsoleApplication1.cpp : Defines the entry point for the console application.
//


// intentionally write the same LDFLAGS differently
#cgo linux LDFLAGS: -L../lib -lhello
#cgo darwin LDFLAGS: -L../lib -lhello

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include "windows.h"

typedef int int32;

// typedef unsigned char uchar;

#define RtlEqualMemory(Destination,Source,Length) (!memcmp((Destination),(Source),(Length)))
#define RtlMoveMemory(Destination,Source,Length) memmove((Destination),(Source),(Length))
#define RtlCopyMemory(Destination,Source,Length) memcpy((Destination),(Source),(Length))
#define RtlFillMemory(Destination,Length,Fill) memset((Destination),(Fill),(Length))
#define RtlZeroMemory(Destination,Length) memset((Destination),0,(Length))

#define LZHCALL
#define LZHEXPORT
#define LZH_DYNAMIC_BUF
#define MALLOC(x) malloc(x)
#define FREE(x)	free(x)

// LZSS Parameters

#define LZH_N			4096	// Size of string buffer
#define LZH_F			60		// Size of look-ahead buffer
#define LZH_THRESHOLD	2
#define LZH_NIL 		LZH_N	// End of tree's node

// Huffman coding parameters

#define LZH_N_CHAR		(256 - LZH_THRESHOLD + LZH_F)
// character code (= 0..LZH_N_CHAR-1)
#define LZH_T		(LZH_N_CHAR * 2 - 1)	// Size of table
#define LZH_R		(LZH_T - 1) 		// root position
#define MAX_FREQ	0x8000
// update when cumulative frequency
// reaches to this value

// Converted from global variables to struct Apr-21-2003
typedef struct {

#ifdef LZH_DYNAMIC_BUF

	unsigned char*	text_buf;
	short int		match_position, match_length,
		*lson, *rson, *dad;

	unsigned short*	freq;	 // cumulative freq table

	//
//////	* pointing parent nodes.
//	* area [LZH_T..(LZH_T + LZH_N_CHAR - 1)] are pointers for leaves

	short int*		prnt;

	// pointing children nodes (son[], son[] + 1)
	short int*		son;

#else	// STATIC

	unsigned char	text_buf[LZH_N + LZH_F - 1];
	short int		match_position, match_length,
		lson[LZH_N + 1], rson[LZH_N + 257], dad[LZH_N + 1];

	unsigned short	freq[LZH_T + 1];   // cumulative freq table
	short int		prnt[LZH_T + LZH_N_CHAR];
	short int		son[LZH_T + 1];		  // bug fixed by Digital Dynamics

#endif

	unsigned short	getbuf;		// Was just "unsigned" fixed 04/12/95
	unsigned char			getlen;
	unsigned		putbuf;
	unsigned char			putlen;

	unsigned short	code, len;

} lzh_t;

static void lzh_init_tree(lzh_t* lzh)  // Initializing tree
{
	short int  i;

	for (i = LZH_N + 1; i <= LZH_N + 256; i++)
		lzh->rson[i] = LZH_NIL;			// root
	for (i = 0; i < LZH_N; i++)
		lzh->dad[i] = LZH_NIL;			// node
}

//****************************
// Inserting node to the tree
// Only used during encoding
//****************************
static void lzh_insert_node(lzh_t* lzh, short int r)
{
	short int  i, p, cmp;
	unsigned char  *key;
	unsigned c;

	cmp = 1;
	key = lzh->text_buf+r;
	p = LZH_N + 1 + key[0];
	lzh->rson[r] = lzh->lson[r] = LZH_NIL;
	lzh->match_length = 0;
	for ( ; ; ) {
		if (cmp >= 0) {
			if (lzh->rson[p] != LZH_NIL)
				p = lzh->rson[p];
			else {
				lzh->rson[p] = r;
				lzh->dad[r] = p;
				return;
			}
		} else {
			if (lzh->lson[p] != LZH_NIL)
				p = lzh->lson[p];
			else {
				lzh->lson[p] = r;
				lzh->dad[r] = p;
				return;
			}
		}
		for (i = 1; i < LZH_F; i++)
			if ((cmp = key[i] - lzh->text_buf[p + i]) != 0)
				break;
		if (i > LZH_THRESHOLD) {
			if (i > lzh->match_length) {
				lzh->match_position = ((r - p) & (LZH_N - 1)) - 1;
				if ((lzh->match_length = i) >= LZH_F)
					break;
			}
			if (i == lzh->match_length) {
				if ((c = ((r - p) & (LZH_N - 1)) - 1)
					< (unsigned)lzh->match_position) {
						lzh->match_position = (short) c;
					}
			}
		}
	}
	lzh->dad[r] = lzh->dad[p];
	lzh->lson[r] = lzh->lson[p];
	lzh->rson[r] = lzh->rson[p];
	lzh->dad[lzh->lson[p]] = r;
	lzh->dad[lzh->rson[p]] = r;
	if (lzh->rson[lzh->dad[p]] == p)
		lzh->rson[lzh->dad[p]] = r;
	else
		lzh->lson[lzh->dad[p]] = r;
	lzh->dad[p] = LZH_NIL;  // remove p
}

static void lzh_delete_node(lzh_t* lzh, short int p)  // Deleting node from the tree
{
	short int  q;

	if (lzh->dad[p] == LZH_NIL)
		return;			// unregistered
	if (lzh->rson[p] == LZH_NIL)
		q = lzh->lson[p];
	else
		if (lzh->lson[p] == LZH_NIL)
			q = lzh->rson[p];
		else {
			q = lzh->lson[p];
			if (lzh->rson[q] != LZH_NIL) {
				do {
					q = lzh->rson[q];
				} while (lzh->rson[q] != LZH_NIL);
				lzh->rson[lzh->dad[q]] = lzh->lson[q];
				lzh->dad[lzh->lson[q]] = lzh->dad[q];
				lzh->lson[q] = lzh->lson[p];
				lzh->dad[lzh->lson[p]] = q;
			}
			lzh->rson[q] = lzh->rson[p];
			lzh->dad[lzh->rson[p]] = q;
		}
		lzh->dad[q] = lzh->dad[p];
		if (lzh->rson[lzh->dad[p]] == p)
			lzh->rson[lzh->dad[p]] = q;
		else
			lzh->lson[lzh->dad[p]] = q;
		lzh->dad[p] = LZH_NIL;
}


// encoder table
static unsigned char lzh_p_len[64] = {
	0x03, 0x04, 0x04, 0x04, 0x05, 0x05, 0x05, 0x05,
		0x05, 0x05, 0x05, 0x05, 0x06, 0x06, 0x06, 0x06,
		0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
		0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08
};

static unsigned char lzh_p_code[64] = {
	0x00, 0x20, 0x30, 0x40, 0x50, 0x58, 0x60, 0x68,
		0x70, 0x78, 0x80, 0x88, 0x90, 0x94, 0x98, 0x9C,
		0xA0, 0xA4, 0xA8, 0xAC, 0xB0, 0xB4, 0xB8, 0xBC,
		0xC0, 0xC2, 0xC4, 0xC6, 0xC8, 0xCA, 0xCC, 0xCE,
		0xD0, 0xD2, 0xD4, 0xD6, 0xD8, 0xDA, 0xDC, 0xDE,
		0xE0, 0xE2, 0xE4, 0xE6, 0xE8, 0xEA, 0xEC, 0xEE,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7,
		0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF
};

// decoder table
static unsigned char lzh_d_code[256] = {
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
		0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
		0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03,
		0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03,
		0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
		0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09,
		0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A,
		0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B,
		0x0C, 0x0C, 0x0C, 0x0C, 0x0D, 0x0D, 0x0D, 0x0D,
		0x0E, 0x0E, 0x0E, 0x0E, 0x0F, 0x0F, 0x0F, 0x0F,
		0x10, 0x10, 0x10, 0x10, 0x11, 0x11, 0x11, 0x11,
		0x12, 0x12, 0x12, 0x12, 0x13, 0x13, 0x13, 0x13,
		0x14, 0x14, 0x14, 0x14, 0x15, 0x15, 0x15, 0x15,
		0x16, 0x16, 0x16, 0x16, 0x17, 0x17, 0x17, 0x17,
		0x18, 0x18, 0x19, 0x19, 0x1A, 0x1A, 0x1B, 0x1B,
		0x1C, 0x1C, 0x1D, 0x1D, 0x1E, 0x1E, 0x1F, 0x1F,
		0x20, 0x20, 0x21, 0x21, 0x22, 0x22, 0x23, 0x23,
		0x24, 0x24, 0x25, 0x25, 0x26, 0x26, 0x27, 0x27,
		0x28, 0x28, 0x29, 0x29, 0x2A, 0x2A, 0x2B, 0x2B,
		0x2C, 0x2C, 0x2D, 0x2D, 0x2E, 0x2E, 0x2F, 0x2F,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
		0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E, 0x3F,
};

static unsigned char lzh_d_len[256] = {
	0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03,
		0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03,
		0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03,
		0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03,
		0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
		0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
		0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
		0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
		0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
		0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
		0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
		0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
		0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
};


static int lzh_getbit(lzh_t* lzh, unsigned char *inbuf, int32 *incnt, int32 inlen)    // get one bit
{
	short int i;

	while (lzh->getlen <= 8) {
		if((*incnt)>=inlen)
			i=0;
		else
			i=inbuf[(*incnt)++];
		lzh->getbuf |= i << (8 - lzh->getlen);
		lzh->getlen += 8;
	}
	i = lzh->getbuf;
	lzh->getbuf <<= 1;
	lzh->getlen--;
	return (i < 0);
}

static short int lzh_getbyte(lzh_t* lzh, unsigned char *inbuf, int32 *incnt, int32 inlen)   // get a byte
{
	unsigned short i;

	while (lzh->getlen <= 8) {
		if((*incnt)>=inlen)
			i=0;
		else
			i=inbuf[(*incnt)++];
		lzh->getbuf |= i << (8 - lzh->getlen);
		lzh->getlen += 8;
	}
	i = lzh->getbuf;
	lzh->getbuf <<= 8;
	lzh->getlen -= 8;
	return i >> 8;
}


// output c bits
static void lzh_putcode(lzh_t* lzh, short int l, unsigned short c, unsigned char *outbuf, int32 *outlen)
{
	lzh->putbuf |= c >> lzh->putlen;
	if ((lzh->putlen += l) >= 8) {
		outbuf[(*outlen)++]=(lzh->putbuf >> 8);
		if ((lzh->putlen -= 8) >= 8) {
			outbuf[(*outlen)++]=(unsigned char)lzh->putbuf;
			lzh->putlen -= 8;
			lzh->putbuf = c << (l - lzh->putlen);
		} else {
			lzh->putbuf <<= 8;
		}
	}
}


// initialize freq tree

static void lzh_start_huff(lzh_t* lzh)
{
	short int i, j;

	for (i = 0; i < LZH_N_CHAR; i++) {
		lzh->freq[i] = 1;
		lzh->son[i] = i + LZH_T;
		lzh->prnt[i + LZH_T] = i;
	}
	i = 0; j = LZH_N_CHAR;
	while (j <= LZH_R) {
		lzh->freq[j] = lzh->freq[i] + lzh->freq[i + 1];
		lzh->son[j] = i;
		lzh->prnt[i] = lzh->prnt[i + 1] = j;
		i += 2; j++;
	}
	lzh->freq[LZH_T] = 0xffff;
	lzh->prnt[LZH_R] = 0;
}


// reconstruct freq tree

static void lzh_reconst(lzh_t* lzh)
{
	short int i, j, k;
	unsigned short f, l;

	// halven cumulative freq for leaf nodes
	j = 0;
	for (i = 0; i < LZH_T; i++) {
		if (lzh->son[i] >= LZH_T) {
			lzh->freq[j] = (lzh->freq[i] + 1) / 2;
			lzh->son[j] = lzh->son[i];
			j++;
		}
	}
	// make a tree : first, connect children nodes
	for (i = 0, j = LZH_N_CHAR; j < LZH_T; i += 2, j++) {
		k = i + 1;
		f = lzh->freq[j] = lzh->freq[i] + lzh->freq[k];
		for (k = j - 1; f < lzh->freq[k]; k--);
		k++;
		l = (j - k) * 2;

		// movmem() is Turbo-C dependent
		//rewritten to memmove() by Kenji

		// movmem(&lzh->freq[k], &lzh->freq[k + 1], l);
		(void)RtlMoveMemory(lzh->freq+k+1,lzh->freq+k, l);
		lzh->freq[k] = f;
		// movmem(&lzh->son[k], &lzh->son[k + 1], l);
		(void)RtlMoveMemory(lzh->son+k+1,lzh->son+k, l);
		lzh->son[k] = i;
	}
	// connect parent nodes
	for (i = 0; i < LZH_T; i++) {
		if ((k = lzh->son[i]) >= LZH_T) {
			lzh->prnt[k] = i;
		} else {
			lzh->prnt[k] = lzh->prnt[k + 1] = i;
		}
	}
}

// update freq tree

static void lzh_update(lzh_t* lzh, short int c)
{
	short int i, j, k, l;

	if (lzh->freq[LZH_R] == MAX_FREQ) {
		lzh_reconst(lzh);
	}
	c = lzh->prnt[c + LZH_T];
	do {
		k = ++lzh->freq[c];

		// swap nodes to keep the tree freq-ordered
		if (k > lzh->freq[l = c + 1]) {
			while (k > lzh->freq[++l]);
			l--;
			lzh->freq[c] = lzh->freq[l];
			lzh->freq[l] = k;

			i = lzh->son[c];
			lzh->prnt[i] = l;
			if (i < LZH_T) lzh->prnt[i + 1] = l;

			j = lzh->son[l];
			lzh->son[l] = i;

			lzh->prnt[j] = c;
			if (j < LZH_T) lzh->prnt[j + 1] = c;
			lzh->son[c] = j;

			c = l;
		}
	} while ((c = lzh->prnt[c]) != 0);	// do it until reaching the root
}

static void lzh_encode_char(lzh_t* lzh, unsigned short c, unsigned char *outbuf, int32 *outlen)
{
	unsigned short i;
	short int j, k;

	i = 0;
	j = 0;
	k = lzh->prnt[c + LZH_T];

	// search connections from leaf node to the root
	do {
		i >>= 1;

		if (k & 1) i += 0x8000;

		j++;
	} while ((k = lzh->prnt[k]) != LZH_R);
	lzh_putcode(lzh, j, i, outbuf, outlen);
	lzh->code = i;
	lzh->len = j;
	lzh_update(lzh,c);
}

static void lzh_encode_position(lzh_t* lzh, unsigned short c, unsigned char *outbuf, int32 *outlen)
{
	unsigned short i;

	// output upper 6 bits with encoding
	i = c >> 6;
	lzh_putcode(lzh, lzh_p_len[i], (unsigned short)(lzh_p_code[i] << 8), outbuf, outlen);

	// output lower 6 bits directly
	lzh_putcode(lzh, 6, (unsigned short)((c & 0x3f) << 10), outbuf, outlen);
}

static void lzh_encode_end(lzh_t* lzh, unsigned char *outbuf, int32 *outlen)
{
	if (lzh->putlen) {
		outbuf[(*outlen)++]=(lzh->putbuf >> 8);
	}
}

static short int lzh_decode_char(lzh_t* lzh, unsigned char *inbuf, int32 *incnt, int32 inlen)
{
	unsigned short c;

	c = lzh->son[LZH_R];


	while (c < LZH_T) {
		c += (unsigned short)lzh_getbit(lzh,inbuf,incnt,inlen);
		c = lzh->son[c];
	}
	c -= LZH_T;
	lzh_update(lzh,c);
	return c;
}

static short int lzh_decode_position(lzh_t* lzh, unsigned char *inbuf, int32 *incnt, int32 inlen)
{
	unsigned short i, j, c;

	// decode upper 6 bits from given table
	i = lzh_getbyte(lzh,inbuf,incnt,inlen);
	c = (unsigned)lzh_d_code[i] << 6;
	j = lzh_d_len[i];

	// input lower 6 bits directly
	j -= 2;
	while (j--) {
		i = (i << 1) + lzh_getbit(lzh,inbuf,incnt,inlen);
	}
	return c | (i & 0x3f);
}

// Compression

// Encoding/Compressing
// Returns length of outbuf
// TODO: Note that inlen usage suggests this is not 64-bit clean
int32 LZHCALL lzh_encode3(unsigned char *inbuf, int32 inlen, unsigned char *outbuf)
{
	strcpy(outbuf,"123456");
	return 6;
}

int32 LZHCALL lzh_encode(unsigned char *inbuf, int32 inlen, unsigned char *outbuf)
{
	MessageBoxA(NULL,inbuf,0,0);
	short int  i, c, len, r, s, last_match_length;
	int32 incnt,outlen; // textsize=0;
	lzh_t lzh;
	RtlZeroMemory(&lzh,sizeof(lzh));

#ifdef LZH_DYNAMIC_BUF

	if((lzh.text_buf=(unsigned char *)MALLOC(LZH_N + LZH_F - 1))==NULL)
		return(-1);
	if((lzh.freq=(unsigned short*)MALLOC((LZH_T + 1)*sizeof(unsigned short)))==NULL) {
		FREE(lzh.text_buf);
		return(-1); }
	if((lzh.prnt=(short *)MALLOC((LZH_T + LZH_N_CHAR)*sizeof(short)))==NULL) {
		FREE(lzh.text_buf);
		FREE(lzh.freq);
		return(-1); }
	if((lzh.son=(short *)MALLOC((LZH_T + 1) * sizeof(short)))==NULL) {
		FREE(lzh.text_buf);
		FREE(lzh.prnt);
		FREE(lzh.freq);
		return(-1); }
	if((lzh.lson=(short *)MALLOC((LZH_N + 1)*sizeof(short)))==NULL) {
		FREE(lzh.text_buf);
		FREE(lzh.prnt);
		FREE(lzh.freq);
		FREE(lzh.son);
		return(-1); }
	if((lzh.rson=(short *)MALLOC((LZH_N + 257)*sizeof(short)))==NULL) {
		FREE(lzh.text_buf);
		FREE(lzh.prnt);
		FREE(lzh.freq);
		FREE(lzh.son);
		FREE(lzh.lson);
		return(-1); }
	if((lzh.dad=(short *)MALLOC((LZH_N + 1)*sizeof(short)))==NULL) {
		FREE(lzh.text_buf);
		FREE(lzh.prnt);
		FREE(lzh.freq);
		FREE(lzh.son);
		FREE(lzh.lson);
		FREE(lzh.rson);
		return(-1); }
#endif

	incnt=0;
	RtlCopyMemory(outbuf,&inlen,sizeof(inlen));
	outlen=sizeof(inlen);
	if(!inlen) {
#ifdef LZH_DYNAMIC_BUF
		FREE(lzh.text_buf);
		FREE(lzh.prnt);
		FREE(lzh.freq);
		FREE(lzh.son);
		FREE(lzh.lson);
		FREE(lzh.rson);
		FREE(lzh.dad);
#endif
		return(outlen); }
	lzh_start_huff(&lzh);
	lzh_init_tree(&lzh);
	s = 0;
	r = LZH_N - LZH_F;
	for (i = s; i < r; i++)
		lzh.text_buf[i] = ' ';
	for (len = 0; len < LZH_F && incnt<inlen; len++)
		lzh.text_buf[r + len] = inbuf[incnt++];
	// textsize = len;
	for (i = 1; i <= LZH_F; i++)
		lzh_insert_node(&lzh,(short)(r - i));
	lzh_insert_node(&lzh,r);
	do {
		if (lzh.match_length > len)
			lzh.match_length = len;
		if (lzh.match_length <= LZH_THRESHOLD) {
			lzh.match_length = 1;
			lzh_encode_char(&lzh,lzh.text_buf[r],outbuf,&outlen);
		} else {
			lzh_encode_char(&lzh,(unsigned short)(255 - LZH_THRESHOLD + lzh.match_length)
				,outbuf,&outlen);
			lzh_encode_position(&lzh,lzh.match_position
				,outbuf,&outlen);
		}
		last_match_length = lzh.match_length;
		for (i = 0; i < last_match_length && incnt<inlen; i++) {
			lzh_delete_node(&lzh,s);
			c=inbuf[incnt++];
			lzh.text_buf[s] = (unsigned char)c;
			if (s < LZH_F - 1)
				lzh.text_buf[s + LZH_N] = (unsigned char)c;
			s = (s + 1) & (LZH_N - 1);
			r = (r + 1) & (LZH_N - 1);
			lzh_insert_node(&lzh,r);
		}

		while (i++ < last_match_length) {
			lzh_delete_node(&lzh,s);
			s = (s + 1) & (LZH_N - 1);
			r = (r + 1) & (LZH_N - 1);
			if (--len) lzh_insert_node(&lzh,r);
		}
	} while (len > 0);
	lzh_encode_end(&lzh,outbuf,&outlen);


#ifdef LZH_DYNAMIC_BUF
	FREE(lzh.text_buf);
	FREE(lzh.prnt);
	FREE(lzh.freq);
	FREE(lzh.son);
	FREE(lzh.lson);
	FREE(lzh.rson);
	FREE(lzh.dad);
#endif

	return(outlen);
}

// Decoding/Uncompressing
// Returns length of outbuf

int32 LZHCALL lzh_decode(unsigned char *inbuf, int32 inlen, unsigned char *outbuf)
{
	short int  i, j, k, r, c;
	unsigned  int  count;
	int32 incnt,textsize;
	lzh_t lzh;

	RtlZeroMemory(&lzh,sizeof(lzh));
#ifdef LZH_DYNAMIC_BUF

	if((lzh.text_buf=(unsigned char *)MALLOC((LZH_N + LZH_F - 1)*2))==NULL)
		return(-1);
	if((lzh.freq=(unsigned short *)MALLOC((LZH_T + 1)*sizeof(unsigned short)))
		==NULL) {
			FREE(lzh.text_buf);
			return(-1); }
		if((lzh.prnt=(short *)MALLOC((LZH_T + LZH_N_CHAR)*sizeof(short)))==NULL) {
			FREE(lzh.text_buf);
			FREE(lzh.freq);
			return(-1); }
		if((lzh.son=(short *)MALLOC((LZH_T + 1) * sizeof(short)))==NULL) {
			FREE(lzh.text_buf);
			FREE(lzh.prnt);
			FREE(lzh.freq);
			return(-1); }

#endif

		incnt=0;
		memcpy(&textsize,inbuf,sizeof(textsize));
		incnt+=sizeof(textsize);
		if (textsize == 0) {
#ifdef LZH_DYNAMIC_BUF
			FREE(lzh.text_buf);
			FREE(lzh.prnt);
			FREE(lzh.freq);
			FREE(lzh.son);
#endif
			return(textsize); }
		lzh_start_huff(&lzh);
		for (i = 0; i < LZH_N - LZH_F; i++)
			*(lzh.text_buf+i) = ' ';
		r = LZH_N - LZH_F;
		for (count = 0; count < (unsigned int)textsize; ) {
			c = lzh_decode_char(&lzh,inbuf,&incnt,inlen);
			if (c < 256) {
				outbuf[count]=(unsigned char)c;
#if 0
				if(r>(LZH_N + LZH_F - 1) || r<0) {
					printf("Overflow! (%d)\n",r);
					getch();
					exit(-1); }
#endif
				*(lzh.text_buf+r) = (unsigned char)c;
				r++;
				r &= (LZH_N - 1);
				count++;
			} else {
				i = (r - lzh_decode_position(&lzh,inbuf,&incnt,inlen) - 1)
					& (LZH_N - 1);
				j = c - 255 + LZH_THRESHOLD;
				for (k = 0; k < j && count<(unsigned int)textsize; k++) {
					c = lzh.text_buf[(i + k) & (LZH_N - 1)];
					outbuf[count]=(unsigned char)c;
#if 0
					if(r>(LZH_N + LZH_F - 1) || r<0) {
						printf("Overflow! (%d)\n",r);
						exit(-1); }
#endif
					*(lzh.text_buf+r) = (unsigned char)c;
					r++;
					r &= (LZH_N - 1);
					count++;
				}
			}
		}

#ifdef LZH_DYNAMIC_BUF
		FREE(lzh.text_buf);
		FREE(lzh.prnt);
		FREE(lzh.freq);
		FREE(lzh.son);
#endif

		return(count);
}

*/
import "C"

import (
	"fmt"
	"unsafe"
)

//int32 LZHCALL lzh_encode(unsigned char *inbuf, int32 inlen, unsigned char *outbuf);
func Lzh(str string) []byte {
	cstr := C.CString(str)

	// box := []string{"xing", "jack", "john"}
	// var buf []*C.unsigned char
	// for i, _ := range bs {
	// 	buf = append(buf, (*C.unsigned char)(unsafe.Pointer(&bs[i])))
	// }
	// box2 := (*C.unsigned char)(unsafe.Pointer(&buf[0]))
	// C.showStringArray(box2)

	// C.String(bs)
	// length := C.lzh_encode(unsigned char *C.unsigned char, inlen C.int, unsigned char *C.unsigned char)
	pRead := C.malloc(C.size_t(len(str) + 1024))
	// length := C.lzh_encode(box2, int32(len(bs)), *C.unsigned char(pRead))
	// C.free(unsafe.Pointer(pRead))
	// return C.GoBytes(pRead, int(length))

	// ch :=

	// cstrPtr := unsafe.Pointer(cstr)

	length := C.lzh_encode((*C.uchar)(unsafe.Pointer(cstr)), C.int32(len(str)), (*C.uchar)(pRead))

	fmt.Println(C.int(length))
	bs := C.GoBytes(pRead, C.int(length))

	C.free(unsafe.Pointer(pRead))
	C.free(unsafe.Pointer(cstr))

	return bs
}

func LzhByBytes(bs []byte) []byte {
	pRead := C.malloc(C.size_t(len(bs) + 1024))

	length := C.lzh_encode((*C.uchar)(unsafe.Pointer(&bs[0])), C.int32(len(bs)), (*C.uchar)(pRead))

	fmt.Println(len(bs))

	fmt.Println(C.int(length))
	bsout := C.GoBytes(pRead, C.int(length))

	C.free(unsafe.Pointer(pRead))
	// C.free(unsafe.Pointer(cstr))

	return bsout
}

func main() {
	bs := []byte("123456789012345678901234567890")

	bsout := LzhByBytes(bs)
	fmt.Println(bsout)
}
